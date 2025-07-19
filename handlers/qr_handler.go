package handlers

import (
	"gopay-clone/models"
	"gopay-clone/services"
	"gopay-clone/utils"
	"gopay-clone/validator"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo"
)

type QRHandler struct {
	qrService *services.QRService
}

func NewQRHandler(qrService *services.QRService) *QRHandler {
	return &QRHandler{qrService: qrService}
}

func (h *QRHandler) CreateQR(c echo.Context) error {
	var req validator.CreateQRRequest
	if err := utils.BindAndValidate(c, &req, validator.ValidateCreateQR); err != nil {
		return err
	}

	qr := &models.QrCode{
		ReceiverAccountID: uint(req.ReceiverAccountID),
		Amount:            req.Amount,
		URL:               req.URL,
		ExpiresAt:         time.Now().Add(time.Minute),
	}

	if err := h.qrService.CreateQR(qr); err != nil {
		if err.Error() == "not found" {
			return utils.NotFoundResponse(c, "receiver account")
		}
		return utils.InternalErrorResponse(c, err)
	}

	return utils.SuccessResponse(c, http.StatusCreated, "qr created successfully", qr)
}

func (h *QRHandler) ScanQr(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("qr_id"))
	if err != nil {
		return utils.ValidationErrorResponse(c, err)
	}
	var req validator.ScanQRRequest
	if err := utils.BindAndValidate(c, &req, validator.ValidateScanQR); err != nil {
		return err
	}

	foundQr, error := h.qrService.GetQRById(uint(id))
	if error != nil {
		return utils.NotFoundResponse(c, "qr")
	}

	if err := h.qrService.ScanQR(foundQr, uint(req.SenderAccountID)); err != nil {
		return utils.ValidationErrorResponse(c, err)
	}
	return utils.SuccessResponse(c, http.StatusOK, "qr scanned", foundQr)
}
