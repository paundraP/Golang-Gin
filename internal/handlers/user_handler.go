package handler

import (
	"net/http"
	"realtime-score/internal/dto"
	"realtime-score/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/sirupsen/logrus"
)

var validate = validator.New()

type UserHandler struct {
	log         *logrus.Logger
	UserService *services.UserService
}

func NewUser(log *logrus.Logger, userService *services.UserService) *UserHandler {
	return &UserHandler{
		log:         log,
		UserService: userService,
	}
}

func (h *UserHandler) Register(c *gin.Context) {
	var req dto.UserCreateRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request", "details": err.Error()})
		return
	}

	if validationErr := validate.Struct(req); validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Validation failed", "details": validationErr.Error()})
		return
	}

	result, err := h.UserService.Register(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Registration successful", "data": result})
}

func (h *UserHandler) Login(c *gin.Context) {
	var req dto.UserLoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := h.UserService.Login(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Login failed",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "login successful",
		"data":    result,
	})
}

func (h *UserHandler) Me(c *gin.Context) {
	user_id := c.MustGet("user_id").(string)

	result, err := h.UserService.GetUserByID(user_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "failed get user",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success get user",
		"data":    result,
	})
}

func (h *UserHandler) GetAllUser(c *gin.Context) {
	result, err := h.UserService.GetAllUser()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "failed get all user",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success get all user",
		"data":    result,
	})
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	user_id := c.MustGet("user_id").(string)
	var req dto.UserUpdateRequest

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request", "details": err.Error()})
		return
	}

	if validationErr := validate.Struct(req); validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Validation failed", "details": validationErr.Error()})
		return
	}

	result, err := h.UserService.UpdateUser(req, user_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "failed update user",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success update user",
		"data":    result,
	})
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	user_id := c.Param("user_id")
	result := h.UserService.DeleteUser(user_id)
	if result != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "failed delete user",
			"message": result.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success delete user",
	})
}
