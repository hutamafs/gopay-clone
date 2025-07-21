package handlers

import (
	"gopay-clone/models"
	"gopay-clone/services"
	"gopay-clone/utils"
	"gopay-clone/validator"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type TransactionHandler struct {
	transactionService *services.TransactionService
}

func NewTransactionHandler(transactionService *services.TransactionService) *TransactionHandler {
	return &TransactionHandler{transactionService: transactionService}
}

func (h *TransactionHandler) CreateTransaction(c echo.Context) error {
	var req validator.CreateTransactionRequest
	if err := utils.BindAndValidate(c, &req, validator.ValidateCreateTransaction); err != nil {
		return err
	}

	transaction := &models.Transaction{
		Amount:            req.Amount,
		SenderAccountId:   req.SenderAccountId,
		ReceiverAccountId: req.ReceiverAccountId,
		Category:          req.Category,
		Type:              req.Type,
	}

	if err := h.transactionService.CreateTransaction(transaction); err != nil {
		if err.Error() == "record not found" {
			return utils.NotFoundResponse(c, "receiver or sender ids")
		}
		if err.Error() == "insufficient balance" {
			return utils.ValidationErrorResponse(c, err)
		}
		return utils.InternalErrorResponse(c, err)
	}

	return utils.SuccessResponse(c, http.StatusCreated, "Transaction created successfully", transaction)
}

func (h *TransactionHandler) GetTransactionDetail(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("transaction_id"))
	if err != nil {
		return utils.ValidationErrorResponse(c, err)
	}

	transaction, err := h.transactionService.GetTransactionById(uint(id))
	if err != nil {
		return utils.NotFoundResponse(c, "transaction id")
	}
	return utils.SuccessResponse(c, http.StatusOK, "Transaction fetched successfully", transaction)
}
