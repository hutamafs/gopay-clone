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

type UserHandler struct {
	userService    *services.UserService
	accountService *services.AccountService
}

func NewUserHandler(userService *services.UserService, accountService *services.AccountService) *UserHandler {
	return &UserHandler{userService: userService, accountService: accountService}
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

func (h *UserHandler) Login(c echo.Context) error {
	var req validator.LoginRequest
	if err := utils.BindAndValidate(c, &req, validator.ValidateLogin); err != nil {
		return err
	}

	user := &models.LoggedinUser{
		Email:    req.Email,
		Password: req.Password,
	}

	token, err := h.userService.Login(user)
	if err != nil {
		return utils.ValidationErrorResponse(c, err)
	}
	// Set the Authorization header with Bearer token
	c.Response().Header().Set("Authorization", "Bearer "+token)

	return utils.SuccessResponse(c, http.StatusOK, "Login successful", map[string]string{"token": token})
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
	loggedInUserId := utils.CLaimJwt(c)

	if id != int(loggedInUserId) {
		return utils.ForbiddenResponse(c, errors.New("unauthorized access"))
	}
	var req validator.UpdateUserRequest
	if err := utils.BindAndValidate(c, &req, validator.ValidateUpdateUser); err != nil {
		return err
	}
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return err
	}
	user.Name = req.Name
	if req.Password != "" {
		user.Password = hashedPassword
	}
	if err := h.userService.UpdateUser(user); err != nil {
		return utils.InternalErrorResponse(c, err)
	}
	return utils.SuccessResponse(c, http.StatusOK, "User updated successfully", user)
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

func (h *UserHandler) GetAccountsByUser(c echo.Context) error {
	userId, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		return utils.ValidationErrorResponse(c, err)
	}
	loggedInUserId := utils.CLaimJwt(c)

	if userId != int(loggedInUserId) {
		return utils.ForbiddenResponse(c, errors.New("unauthorized access"))
	}

	accounts, err := h.accountService.GetAccountsByUser(uint(userId))
	if err != nil {
		return utils.NotFoundResponse(c, "user id")
	}
	return utils.SuccessResponse(c, http.StatusOK, "Accounts for user fetched successfully", accounts)
}
