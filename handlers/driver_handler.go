package handlers

import (
	"gopay-clone/models"
	"gopay-clone/services"
	"gopay-clone/utils"
	"gopay-clone/validator"
	"net/http"

	"github.com/labstack/echo/v4"
)

type DriverHandler struct {
	driverService *services.DriverService
	userService   *services.UserService
}

func NewDriverHandler(driverService *services.DriverService, userService *services.UserService) *DriverHandler {
	return &DriverHandler{driverService: driverService, userService: userService}
}

func (h *DriverHandler) CreateDriver(c echo.Context) error {
	var req validator.CreateDriverRequest
	if err := utils.BindAndValidate(c, &req, validator.ValidateCreateDriver); err != nil {
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

	driver := &models.DriverProfile{
		UserId:            user.ID,
		LicenseNumber:     req.LicenseNumber,
		LicensePictureURL: req.LicensePictureURL,
		VehiclePlate:      req.VehiclePlate,
		VehicleType:       req.VehicleType,
		CurrentLocation:   req.CurrentLocation,
	}

	if err := h.driverService.CreateDriverProfile(driver); err != nil {
		return utils.InternalErrorResponse(c, err)
	}

	return utils.SuccessResponse(c, http.StatusCreated, "Driver created successfully", driver)
}

func (h *DriverHandler) GetDriverByID(c echo.Context) error {
	loggedInUserId := utils.CLaimJwt(c)

	driver, err := h.driverService.GetDriverByUserID(uint(loggedInUserId))
	if err != nil {
		return utils.NotFoundResponse(c, "driver profile")
	}

	return utils.SuccessResponse(c, http.StatusOK, "Driver profile fetched successfully", driver)
}

func (h *DriverHandler) GetAllDrivers(c echo.Context) error {
	drivers, err := h.driverService.GetAllDrivers()
	if err != nil {
		return utils.InternalErrorResponse(c, err)
	}

	return utils.SuccessResponse(c, http.StatusOK, "Drivers fetched successfully", drivers)
}

func (h *DriverHandler) GetAvailableDrivers(c echo.Context) error {
	vehicleType := c.QueryParam("vehicle_type")

	var drivers []models.DriverProfile
	var err error

	if vehicleType != "" {
		drivers, err = h.driverService.GetDriversByVehicleType(models.VehicleType(vehicleType))
	} else {
		drivers, err = h.driverService.GetAvailableDrivers()
	}

	if err != nil {
		return utils.InternalErrorResponse(c, err)
	}

	return utils.SuccessResponse(c, http.StatusOK, "Available drivers fetched successfully", drivers)
}

func (h *DriverHandler) UpdateDriverProfile(c echo.Context) error {
	loggedInUserId := utils.CLaimJwt(c)

	// Get driver profile to verify ownership
	driver, err := h.driverService.GetDriverByUserID(uint(loggedInUserId))
	if err != nil {
		return utils.NotFoundResponse(c, "driver profile")
	}

	var req validator.UpdateDriverRequest
	if err := utils.BindAndValidate(c, &req, validator.ValidateUpdateDriver); err != nil {
		return err
	}

	updates := map[string]any{}
	if req.LicenseNumber != "" {
		updates["license_number"] = req.LicenseNumber
	}
	if req.LicensePictureURL != "" {
		updates["license_picture_url"] = req.LicensePictureURL
	}
	if req.VehiclePlate != "" {
		updates["vehicle_plate"] = req.VehiclePlate
	}
	if req.VehicleType != "" {
		updates["vehicle_type"] = req.VehicleType
	}
	if req.CurrentLocation != "" {
		updates["current_location"] = req.CurrentLocation
	}

	if err := h.driverService.UpdateDriver(driver.ID, updates); err != nil {
		return utils.InternalErrorResponse(c, err)
	}

	return utils.SuccessResponse(c, http.StatusOK, "Driver profile updated successfully", nil)
}

func (h *DriverHandler) UpdateDriverStatus(c echo.Context) error {
	loggedInUserId := utils.CLaimJwt(c)

	var req validator.UpdateDriverStatusRequest
	if err := utils.BindAndValidate(c, &req, validator.ValidateUpdateDriverStatus); err != nil {
		return err
	}

	if err := h.driverService.UpdateDriverStatus(uint(loggedInUserId), req.Status); err != nil {
		return utils.InternalErrorResponse(c, err)
	}

	return utils.SuccessResponse(c, http.StatusOK, "Driver status updated successfully", nil)
}

func (h *DriverHandler) UpdateDriverLocation(c echo.Context) error {
	loggedInUserId := utils.CLaimJwt(c)

	var req validator.UpdateDriverLocationRequest
	if err := utils.BindAndValidate(c, &req, validator.ValidateUpdateDriverLocation); err != nil {
		return err
	}

	if err := h.driverService.UpdateDriverLocation(uint(loggedInUserId), req.CurrentLocation); err != nil {
		return utils.InternalErrorResponse(c, err)
	}

	return utils.SuccessResponse(c, http.StatusOK, "Driver location updated successfully", nil)
}

func (h *DriverHandler) DeleteDriverProfile(c echo.Context) error {
	loggedInUserId := utils.CLaimJwt(c)

	// Get driver profile to verify ownership and get ID
	driver, err := h.driverService.GetDriverByUserID(uint(loggedInUserId))
	if err != nil {
		return utils.NotFoundResponse(c, "driver profile")
	}

	if err := h.driverService.DeleteDriver(driver.ID); err != nil {
		return utils.InternalErrorResponse(c, err)
	}

	return utils.SuccessResponse(c, http.StatusOK, "Driver profile deleted successfully", nil)
}
