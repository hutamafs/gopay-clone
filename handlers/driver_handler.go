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

// CreateDriver godoc
// @Summary Create a new driver
// @Description Register a new driver account
// @Tags Driver
// @Accept json
// @Produce json
// @Param driver body validator.CreateDriverRequest true "Driver info"
// @Success 201 {object} utils.APISuccessResponse{data=models.DriverProfile}
// @Failure 400 {object} utils.APIErrorResponse{error=utils.ErrorDetail}
// @Router /public/drivers [post]
func (h *DriverHandler) CreateDriver(c echo.Context) error {
	var req validator.CreateDriverRequest
	if err := utils.BindAndValidate(c, &req, validator.ValidateCreateDriver); err != nil {
		return err
	}
	hashPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return utils.SplitErrorResponse(c, err)
	}

	user := &models.User{
		Name:              req.Name,
		Email:             req.Email,
		Password:          hashPassword,
		Phone:             req.Phone,
		ProfilePictureURL: req.ProfilePictureURL,
		Type:              "driver",
	}

	if err := h.userService.CreateUser(user); err != nil {
		return utils.SplitErrorResponse(c, err)
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
		return utils.SplitErrorResponse(c, err)
	}

	return utils.SuccessResponse(c, http.StatusCreated, "Driver created successfully", driver)
}

// GetDriverByID retrieves driver profile by user ID from JWT token
// @Summary Get driver profile
// @Description Get driver profile information for the authenticated user
// @Tags drivers
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} utils.APISuccessResponse{data=models.DriverProfile}
// @Failure 400 {object}  utils.APIErrorResponse{error=utils.ErrorDetail}
// @Failure 401 {object}  utils.APIErrorResponse{error=utils.ErrorAuth}
// @Failure 404  {object} utils.APIErrorResponse{error=utils.ErrorNotFound}
// @Failure 500 {object}  utils.APIErrorResponse{error=utils.ErrorDetail}
// @Router /drivers/profile [get]
func (h *DriverHandler) GetDriverByID(c echo.Context) error {
	loggedInUserId := utils.CLaimJwt(c)

	driver, err := h.driverService.GetDriverByUserID(uint(loggedInUserId))
	if err != nil {
		return utils.SplitErrorResponse(c, err)
	}

	return utils.SuccessResponse(c, http.StatusOK, "Driver profile fetched successfully", driver)
}

// GetAllDrivers godoc
// @Summary List all drivers
// @Description Retrieve all registered drivers
// @Tags Driver
// @Produce json
// @Success 200 {object} utils.APISuccessResponse{data=[]models.DriverProfile}
// @Failure 400 {object} utils.APIErrorResponse{error=utils.ErrorDetail}
// @Router /public/drivers [get]
func (h *DriverHandler) GetAllDrivers(c echo.Context) error {
	drivers, err := h.driverService.GetAllDrivers()
	if err != nil {
		return utils.SplitErrorResponse(c, err)
	}

	return utils.SuccessResponse(c, http.StatusOK, "Drivers fetched successfully", drivers)
}

// GetAvailableDrivers godoc
// @Summary List all available drivers
// @Description Retrieve all available drivers
// @Tags Driver
// @Produce json
// @Success 200 {object} utils.APISuccessResponse{data=[]models.DriverProfile}
//
//	@Failure 400 {object} utils.APIErrorResponse{error=utils.ErrorDetail}
//
// @Router /drivers/available [get]
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
		return utils.SplitErrorResponse(c, err)
	}

	return utils.SuccessResponse(c, http.StatusOK, "Available drivers fetched successfully", drivers)
}

// UpdateDriver godoc
// @Summary Update driver by ID
// @Description Update details of a driver
// @Tags Driver
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param driver body validator.UpdateDriverRequest true "Updated driver profile"
// @Success 200 {object} utils.APISuccessResponse{data=models.DriverProfile}
// @Failure 400 {object} utils.APIErrorResponse{error=utils.ErrorDetail}
// @Failure 404 {object} utils.APIErrorResponse{error=utils.ErrorNotFound}
// @Router /drivers/profile [put]
func (h *DriverHandler) UpdateDriverProfile(c echo.Context) error {
	loggedInUserId := utils.CLaimJwt(c)

	// Get driver profile to verify ownership
	driver, err := h.driverService.GetDriverByUserID(uint(loggedInUserId))
	if err != nil {
		return utils.SplitErrorResponse(c, err)
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
		return utils.SplitErrorResponse(c, err)
	}

	return utils.SuccessResponse(c, http.StatusOK, "Driver profile updated successfully", nil)
}

// UpdateDriverStatus godoc
// @Summary Update driver status by ID
// @Description Update status of a driver
// @Tags Driver
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param driver body validator.UpdateDriverStatusRequest true "Updated driver status"
// @Success 200 {object} utils.APISuccessResponse{data=models.DriverProfile}
// @Failure 400 {object} utils.APIErrorResponse{error=utils.ErrorDetail}
// @Failure 404 {object} utils.APIErrorResponse{error=utils.ErrorNotFound}
// @Router /drivers/status [put]
func (h *DriverHandler) UpdateDriverStatus(c echo.Context) error {
	loggedInUserId := utils.CLaimJwt(c)

	var req validator.UpdateDriverStatusRequest
	if err := utils.BindAndValidate(c, &req, validator.ValidateUpdateDriverStatus); err != nil {
		return err
	}
	driver, err := h.driverService.GetDriverByUserID(uint(loggedInUserId))
	if err != nil {
		return utils.SplitErrorResponse(c, err)
	}

	if err := h.driverService.UpdateDriverStatus(uint(driver.ID), req.Status); err != nil {
		return utils.SplitErrorResponse(c, err)
	}

	return utils.SuccessResponse(c, http.StatusOK, "Driver status updated successfully", nil)
}

// UpdateDriverLocation godoc
// @Summary Update driver location by ID
// @Description Update location of a driver
// @Tags Driver
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param driver body validator.UpdateDriverLocationRequest true "Updated driver current location"
// @Success 200 {object} utils.APISuccessResponse{data=models.DriverProfile}
// @Failure 400 {object} utils.APIErrorResponse{error=utils.ErrorDetail}
// @Router /drivers/location [put]
func (h *DriverHandler) UpdateDriverLocation(c echo.Context) error {
	loggedInUserId := utils.CLaimJwt(c)

	var req validator.UpdateDriverLocationRequest
	if err := utils.BindAndValidate(c, &req, validator.ValidateUpdateDriverLocation); err != nil {
		return err
	}

	if err := h.driverService.UpdateDriverLocation(uint(loggedInUserId), req.CurrentLocation); err != nil {
		return utils.SplitErrorResponse(c, err)
	}

	return utils.SuccessResponse(c, http.StatusOK, "Driver location updated successfully", nil)
}

// DeleteDriverProfile godoc
// @Summary Delete Driver profile by ID
// @Description delete driver profile
// @Tags Driver
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} utils.APISuccessResponse{data=models.DriverProfile}
// @Failure 400 {object} utils.APIErrorResponse{error=utils.ErrorDetail}
// @Failure 404 {object} utils.APIErrorResponse{error=utils.ErrorNotFound}

// @Router /drivers/profile [delete]
func (h *DriverHandler) DeleteDriverProfile(c echo.Context) error {
	loggedInUserId := utils.CLaimJwt(c)

	// Get driver profile to verify ownership and get ID
	driver, err := h.driverService.GetDriverByUserID(uint(loggedInUserId))
	if err != nil {
		return utils.SplitErrorResponse(c, err)
	}

	if err := h.driverService.DeleteDriver(driver.ID); err != nil {
		return utils.SplitErrorResponse(c, err)
	}

	return utils.SuccessResponse(c, http.StatusOK, "Driver profile deleted successfully", nil)
}
