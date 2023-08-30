package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jutionck/golang-upskilling-agt/model"
	"github.com/jutionck/golang-upskilling-agt/usecase"
)

type AuthController struct {
	router *gin.Engine
	uc     usecase.AuthUseCase
}

func (a *AuthController) loginHandler(c *gin.Context) {
	var payload model.User
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	token, err := a.uc.Login(payload.Username, payload.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": token})
}

func NewAuthController(r *gin.Engine, uc usecase.AuthUseCase) *AuthController {
	controller := AuthController{
		router: r,
		uc:     uc,
	}
	rg := r.Group("/api/v1")
	rg.POST("/auth/login", controller.loginHandler)
	return &controller
}
