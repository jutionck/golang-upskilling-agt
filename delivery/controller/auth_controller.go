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

// Login godoc
// @Summary Login a user
// @Description Login a user
// @Accept json
// @Produce json
// @Tags Auth
// @Param Body body dto.Auth true "Login"
// @Success 201 {object} string
// @Router /auth/login [post]
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

func (a *AuthController) Route() {
	rg := a.router.Group("/api/v1")
	rg.POST("/auth/login", a.loginHandler)
}

func NewAuthController(r *gin.Engine, uc usecase.AuthUseCase) *AuthController {
	return &AuthController{router: r, uc: uc}
}
