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

type UserHandler struct {
	userService *services.UserService
}

func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) CreateUser(c echo.Context) error {
	var req validator.CreateUserRequest
	if err := utils.BindAndValidate(c, &req, validator.ValidateCreateUser); err != nil {
		return err
	}
	hashPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return err
	}

	user := &models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashPassword,
	}

	if err := h.userService.CreateUser(user); err != nil {
		return utils.InternalErrorResponse(c, err)
	}

	return utils.SuccessResponse(c, http.StatusCreated, "User created successfully", user)
}

func (h *UserHandler) GetUserById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return utils.ValidationErrorResponse(c, err)
	}

	user, err := h.userService.GetUserById(uint(id))
	if err != nil {
		return utils.NotFoundResponse(c, "id")
	}
	return utils.SuccessResponse(c, http.StatusOK, "User fetched successfully", user)
}

func (h *UserHandler) GetAllUsers(c echo.Context) error {
	users, err := h.userService.GetUsers()
	if err != nil {
		return utils.ValidationErrorResponse(c, err)
	}
	return utils.SuccessResponse(c, http.StatusOK, "Users fetched successfully, users", users)
}

func (h *UserHandler) UpdateUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return utils.ValidationErrorResponse(c, err)
	}
	user, err := h.userService.GetUserById(uint(id))
	if err != nil {
		return utils.NotFoundResponse(c, "id")
	}
	var req validator.UpdateUserRequest
	utils.BindAndValidate(c, &req, validator.ValidateUpdateUser)
	updatedUser := &models.User{
		Name:     req.Name,
		Email:    user.Email,
		Password: req.Password,
	}
	updatedUser.ID = uint(id)
	if err := h.userService.UpdateUser(updatedUser); err != nil {
		return utils.InternalErrorResponse(c, err)
	}
	return utils.SuccessResponse(c, http.StatusOK, "User updated successfully", updatedUser)
}

func (h *UserHandler) DeleteUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return utils.ValidationErrorResponse(c, err)
	}
	if err := h.userService.DeleteUser(uint(id)); err != nil {
		return utils.InternalErrorResponse(c, err)
	}
	return utils.SuccessResponse(c, http.StatusOK, "user deleted", nil)
}
