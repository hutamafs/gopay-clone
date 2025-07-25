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
