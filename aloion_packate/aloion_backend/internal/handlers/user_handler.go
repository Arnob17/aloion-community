package handlers

import (
	"fmt"
	"net/http"

	"github.com/arnob17/aloion_backend/internal/dto"
	"github.com/arnob17/aloion_backend/internal/services"
	"github.com/arnob17/aloion_backend/internal/utils"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService *services.UserService
	config      interface {
		GetJWTSecret() string
		GetJWTExpiresIn() int
	}
}

func NewUserHandler(userService *services.UserService, jwtSecret string, jwtExpiresIn int) *UserHandler {
	return &UserHandler{
		userService: userService,
		config: &jwtConfig{
			secret:    jwtSecret,
			expiresIn: jwtExpiresIn,
		},
	}
}

type jwtConfig struct {
	secret    string
	expiresIn int
}

func (c *jwtConfig) GetJWTSecret() string {
	return c.secret
}

func (c *jwtConfig) GetJWTExpiresIn() int {
	return c.expiresIn
}

func (h *UserHandler) Register(c *gin.Context) {
	var req dto.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.userService.Register(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Generate token
	token, err := utils.GenerateToken(user.ID, user.Email, string(user.Role), h.config.GetJWTExpiresIn())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusCreated, dto.AuthResponse{
		User:  dto.ToUserResponse(user),
		Token: token,
	})
}

func (h *UserHandler) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.userService.Login(req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// Generate token
	token, err := utils.GenerateToken(user.ID, user.Email, string(user.Role), h.config.GetJWTExpiresIn())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, dto.AuthResponse{
		User:  dto.ToUserResponse(user),
		Token: token,
	})
}

func (h *UserHandler) GetProfile(c *gin.Context) {
	userID, _ := c.Get("user_id")
	user, err := h.userService.GetByID(userID.(uint))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, dto.ToUserResponse(user))
}

func (h *UserHandler) GetAllUsers(c *gin.Context) {
	limit := 20
	offset := 0

	if l := c.Query("limit"); l != "" {
		if parsed, err := parseInt(l); err == nil {
			limit = parsed
		}
	}
	if o := c.Query("offset"); o != "" {
		if parsed, err := parseInt(o); err == nil {
			offset = parsed
		}
	}

	users, total, err := h.userService.GetAll(limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	responses := make([]dto.UserResponse, len(users))
	for i, user := range users {
		responses[i] = dto.ToUserResponse(&user)
	}

	c.JSON(http.StatusOK, gin.H{
		"users": responses,
		"total": total,
		"limit": limit,
		"offset": offset,
	})
}

func parseInt(s string) (int, error) {
	var result int
	_, err := fmt.Sscanf(s, "%d", &result)
	return result, err
}
