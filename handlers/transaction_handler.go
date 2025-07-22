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
		SenderAccountID:   req.SenderAccountID,
		ReceiverAccountID: req.ReceiverAccountID,
		QrCodeID:          req.QrCodeID,
		ServiceID:         req.ServiceID,
	}

	if req.Type != nil {
		transaction.Type = models.TransactionType(*req.Type)
	}
	if req.Category != nil {
		transaction.Category = models.TransactionCategory(*req.Category)
	}
	if req.Status != nil {
		transaction.Status = models.TransactionStatus(*req.Status)
	}
	if req.ServiceType != nil {
		transaction.ServiceType = models.ServiceType(*req.ServiceType)
	}
	if req.Description != nil {
		transaction.Description = *req.Description
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

func (h *TransactionHandler) UpdateTransactionDetail(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("transaction_id"))
	if err != nil {
		return utils.ValidationErrorResponse(c, err)
	}
	loggedInUserId := utils.CLaimJwt(c)

	if loggedInUserId == 0 {
		return utils.ForbiddenResponse(c, errors.New("unauthorized access"))
	}
	var req validator.UpdateTransactionRequest
	if err := utils.BindAndValidate(c, &req, validator.ValidateUpdateTransaction); err != nil {
		return err
	}
	updates := make(map[string]any)

	if req.Status != nil {
		updates["status"] = *req.Status
	}
	if req.QrCodeID != nil {
		updates["qr_code_id"] = *req.QrCodeID
	}
	if req.Category != nil {
		updates["category"] = *req.Category
	}
	if req.Description != nil {
		updates["description"] = *req.Description
	}

	if err := h.transactionService.UpdateTransaction(uint(id), updates); err != nil {
		return utils.InternalErrorResponse(c, err)
	}
	return utils.SuccessResponse(c, http.StatusOK, "Transaction updated successfully", updates)
}
