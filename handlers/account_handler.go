package handlers

import (
	"gopay-clone/models"
	"gopay-clone/services"
	"gopay-clone/utils"
	"gopay-clone/validator"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type AccountHandler struct {
	accountService *services.AccountService
}

func NewAccountHandler(accountService *services.AccountService) *AccountHandler {
	return &AccountHandler{accountService: accountService}
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
		return utils.InternalErrorResponse(c, err)
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
		return utils.ValidationErrorResponse(c, err)
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
		return utils.ValidationErrorResponse(c, err)
	}
	var req validator.UpdateAccountRequest
	if err := utils.BindAndValidate(c, &req, validator.ValidateUpdateAccount); err != nil {
		return utils.ValidationErrorResponse(c, err)
	}
	updateAccount := &models.Account{
		Name:    req.Name,
		Balance: req.Balance,
		UserId:  account.UserId,
	}
	updateAccount.ID = uint(accountId)
	if err := h.accountService.UpdateAccount(updateAccount); err != nil {
		return utils.InternalErrorResponse(c, err)
	}
	return utils.SuccessResponse(c, http.StatusOK, "Account updated successfully", updateAccount)
}

func (h *AccountHandler) GetAccountDetail(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("account_id"))
	if err != nil {
		return utils.ValidationErrorResponse(c, err)
	}

	user, err := h.accountService.GetAccountById(uint(id))
	if err != nil {
		return utils.NotFoundResponse(c, "id")
	}
	return utils.SuccessResponse(c, http.StatusOK, "Account detail fetched successfully", user)
}
