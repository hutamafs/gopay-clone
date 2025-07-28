package handlers

import (
	"errors"
	"gopay-clone/models"
	"gopay-clone/services"

	"gopay-clone/utils"
	"gopay-clone/validator"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type AccountHandler struct {
	accountService     *services.AccountService
	transactionService *services.TransactionService
}

/*
	return err if its from bind and validate because it will return validation error from the internal function

	return validationerror if the error might comes from user input

	return split error if it might comes from the service
*/

func NewAccountHandler(accountService *services.AccountService, transactionService *services.TransactionService) *AccountHandler {
	return &AccountHandler{accountService: accountService, transactionService: transactionService}
}

// CreateAccount godoc
// @Summary Create a new wallet account
// @Description Register a new wallet account
// @Tags Account
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param account body validator.CreateAccountRequest true "Created account"
// @Success 201 {object} models.Account
// @Failure 400 {object} map[string]interface{}
// @Router /accounts [post]
func (h *AccountHandler) CreateAccount(c echo.Context) error {
	var req validator.CreateAccountRequest
	if err := utils.BindAndValidate(c, &req, validator.ValidateCreateAccount); err != nil {
		return err
	}

	account := &models.Account{
		Name:    req.Name,
		Balance: req.Balance,
		UserId:  req.UserId,
	}

	if err := h.accountService.CreateAccount(account); err != nil {
		return utils.SplitErrorResponse(c, err)
	}

	return utils.SuccessResponse(c, http.StatusCreated, "Account created successfully", account)
}

// GetBalanceByAccount godoc
// @Summary Get latest available balance from an account
// @Description Retrieve wallet balance
// @Param account_id path int true "Account ID"
// @Tags Account
// @Produce json
// @Security BearerAuth
// @Success 200 {object} models.Account
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /accounts [get]
func (h *AccountHandler) GetBalanceByAccountId(c echo.Context) error {
	accountId, err := strconv.Atoi(c.Param("account_id"))
	if err != nil {
		return utils.ValidationErrorResponse(c, err)
	}
	balance, err := h.accountService.GetBalanceByAccountId(uint(accountId))
	if err != nil {
		return utils.SplitErrorResponse(c, err)
	}
	val := map[string]float64{
		"balance": *balance,
	}
	return utils.SuccessResponse(c, http.StatusOK, "Balance for account fetched successfully", val)
}

// UpdateAccountDetail godoc
// @Summary Update account detail e.g. name
// @Description Update wallet account in this case the name
// @Param account_id path int true "Account ID"
// @Tags Account
// @Produce json
// @Security BearerAuth
// @Param account body validator.UpdateAccountRequest true "Updated account name"
// @Success 200 {object} models.Account
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 403 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /accounts/:account_id [put]
func (h *AccountHandler) UpdateAccount(c echo.Context) error {
	accountId, err := strconv.Atoi(c.Param("account_id"))
	if err != nil {
		return utils.ValidationErrorResponse(c, err)
	}
	account, err := h.accountService.GetAccountById(uint(accountId))
	if err != nil {
		return utils.SplitErrorResponse(c, err)
	}
	loggedInUserId := utils.CLaimJwt(c)
	if int(account.UserId) != int(loggedInUserId) {
		return utils.ForbiddenResponse(c, errors.New("unauthorized access"))
	}
	var req validator.UpdateAccountRequest
	if err := utils.BindAndValidate(c, &req, validator.ValidateUpdateAccount); err != nil {
		return err
	}
	account.Name = req.Name
	if err := h.accountService.UpdateAccount(account); err != nil {
		return utils.SplitErrorResponse(c, err)
	}
	return utils.SuccessResponse(c, http.StatusOK, "Account updated successfully", account)
}

// GetAccountDetail godoc
// @Summary Get account detail
// @Description Get Account detail, the name, balance and user id
// @Param account_id path int true "Account ID"
// @Tags Account
// @Produce json
// @Security BearerAuth
// @Success 200 {object} models.Account
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 403 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /accounts/:account_id/detail [get]
func (h *AccountHandler) GetAccountDetail(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("account_id"))
	if err != nil {
		return utils.ValidationErrorResponse(c, err)
	}

	account, err := h.accountService.GetAccountById(uint(id))
	if err != nil {
		return utils.SplitErrorResponse(c, err)
	}
	loggedInUserId := utils.CLaimJwt(c)
	if int(account.UserId) != int(loggedInUserId) {
		return utils.ForbiddenResponse(c, errors.New("unauthorized access"))
	}
	return utils.SuccessResponse(c, http.StatusOK, "Account detail fetched successfully", account)
}

// GetTransactionByAccounts godoc
// @Summary Get latest transactions from account
// @Description Get all latest. transaction from single account
// @Param account_id path int true "Account ID"
// @Tags Account
// @Produce json
// @Security BearerAuth
// @Success 200 {array} models.Account
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /accounts/:account_id/transactions [get]
func (h *AccountHandler) GetTransactionByAccounts(c echo.Context) error {
	accountId, err := strconv.Atoi(c.Param("account_id"))
	if err != nil {
		return utils.ValidationErrorResponse(c, err)
	}

	transactions, err := h.transactionService.GetTransactionsByAccount(uint(accountId))
	if err != nil {
		return utils.NotFoundResponse(c, "user id")
	}
	return utils.SuccessResponse(c, http.StatusOK, "Transactions for account fetched successfully", transactions)
}
