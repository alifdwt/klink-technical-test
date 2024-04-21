package handler

import (
	"net/http"
	"technical-test/domain/requests/auth"

	"github.com/gin-gonic/gin"
)

func (h *Handler) initAuthGroup(api *gin.Engine) {
	auth := api.Group("/api/auth")

	auth.POST("/register", h.handlerRegister)
	auth.POST("/login", h.handlerLogin)
}

// handleRegister
// @Summary Register to the application
// @Description Register a new user
// @Tags auth
// @Accept json
// @Produce json
// @Param data body auth.RegisterRequest true "data"
// @Success 200 {object} responses.Token
// @Failure 400 {object} responses.ErrorMessage
// @Failure 500 {object} responses.ErrorMessage
// @Router /auth/register [post]
func (h *Handler) handlerRegister(c *gin.Context) {
	var body auth.RegisterRequest

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := body.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	res, err := h.services.Auth.Register(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	c.JSON(http.StatusCreated, res)
}

// handleLogin
// @Summary Login to the application
// @Description Login a user
// @Tags auth
// @Accept json
// @Produce json
// @Param data body auth.LoginRequest true "data"
// @Success 200 {object} responses.Token
// @Failure 400 {object} responses.ErrorMessage
// @Failure 500 {object} responses.ErrorMessage
// @Router /auth/login [post]
func (h *Handler) handlerLogin(c *gin.Context) {
	var body auth.LoginRequest

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := body.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	res, err := h.services.Auth.Login(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, res)
}
