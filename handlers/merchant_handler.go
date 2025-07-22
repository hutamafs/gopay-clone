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

type MerchantHandler struct {
	userService     *services.UserService
	merchantService *services.MerchantService
}

func NewMerchantHandler(userService *services.UserService, merchantService *services.MerchantService) *MerchantHandler {
	return &MerchantHandler{userService: userService, merchantService: merchantService}
}

func (h *MerchantHandler) CreateMerchant(c echo.Context) error {
	var req validator.CreateMerchantRequest
	if err := utils.BindAndValidate(c, &req, validator.ValidateCreateMerchant); err != nil {
		return err
	}
	hashPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return err
	}

	user := &models.User{
		Name:              req.Name,
		Email:             req.Email,
		Password:          hashPassword,
		Phone:             req.Phone,
		ProfilePictureURL: req.ProfilePictureURL,
		Type:              "merchant",
	}

	if err := h.userService.CreateUser(user); err != nil {
		return utils.InternalErrorResponse(c, err)
	}
	var merchantPhone = user.Phone

	if req.MerchantPhone != nil {
		merchantPhone = *req.MerchantPhone
	}

	merchant := &models.MerchantProfile{
		UserId:          user.ID,
		Location:        req.Location,
		MerchantName:    req.MerchantName,
		Description:     req.Description,
		MerchantPhone:   merchantPhone,
		Category:        req.Category,
		OpenHour:        req.OpenHour,
		ClosedHour:      req.ClosedHour,
		MerchantLogoURL: req.MerchantLogoURL,
	}
	if err := h.merchantService.CreateMerchant(merchant); err != nil {
		return utils.InternalErrorResponse(c, err)
	}

	return utils.SuccessResponse(c, http.StatusCreated, "Merchant created successfully", merchant)
}

func (h *MerchantHandler) GetAllMerchants(c echo.Context) error {
	merchants, err := h.merchantService.GetAllMerchants()
	if err != nil {
		return utils.ValidationErrorResponse(c, err)
	}
	return utils.SuccessResponse(c, http.StatusOK, "Merchants fetched successfully", merchants)
}

func (h *MerchantHandler) GetMerchantByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("merchant_id"))
	if err != nil {
		return utils.ValidationErrorResponse(c, err)
	}

	merchant, err := h.merchantService.GetMerchantByID(uint(id))
	if err != nil {
		return utils.NotFoundResponse(c, "id")
	}
	return utils.SuccessResponse(c, http.StatusOK, "Merchant Profile fetched successfully", merchant)
}

func (h *MerchantHandler) UpdateMerchantByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("merchant_id"))
	if err != nil {
		return utils.ValidationErrorResponse(c, err)
	}
	merchant, err := h.merchantService.GetMerchantByID(uint(id))
	if err != nil {
		return utils.NotFoundResponse(c, "id")
	}
	loggedInUserId := utils.CLaimJwt(c)

	if merchant.UserId != uint(loggedInUserId) {
		return utils.ForbiddenResponse(c, errors.New("unauthorized access"))
	}
	var req validator.UpdateMerchantRequest
	if err := utils.BindAndValidate(c, &req, validator.ValidateUpdateMerchant); err != nil {
		return err
	}

	updatedMerchant := map[string]any{
		"UserId":          merchant.UserId,
		"Location":        req.Location,
		"MerchantName":    req.MerchantName,
		"Description":     req.Description,
		"MerchantPhone":   req.MerchantPhone,
		"Category":        req.Category,
		"OpenHour":        req.OpenHour,
		"ClosedHour":      req.ClosedHour,
		"MerchantLogoURL": req.MerchantLogoURL,
	}

	if err := h.merchantService.UpdateMerchant(uint(id), updatedMerchant); err != nil {
		return utils.InternalErrorResponse(c, err)
	}
	return utils.SuccessResponse(c, http.StatusOK, "Merchant profile updated successfully", merchant)
}
