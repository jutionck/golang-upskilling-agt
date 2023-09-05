package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jutionck/golang-upskilling-agt/delivery/middleware"
	"github.com/jutionck/golang-upskilling-agt/model"
	"github.com/jutionck/golang-upskilling-agt/usecase"
)

type UserController struct {
	uc             usecase.UserUseCase
	router         *gin.Engine
	authMiddleware middleware.AuthMiddleware
}

// User godoc
// @Summary Create a user
// @Description Creare a new user
// @Accept json
// @Produce json
// @Tags User
// @Param Body body dto.UserRequestDto true "New User"
// @Success 201 {object} model.User
// @Router /users [post]
func (u *UserController) createHandler(c *gin.Context) {
	var payload model.User
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err := u.uc.RegisterNewUser(&payload); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, payload)
}

// ListAccounts godoc
// @Summary      List users
// @Description  get users
// @Tags         User
// @Accept       json
// @Produce      json
// @Security 		 Bearer
// @Success      200  {array}   model.User
// @Router       /users [get]
func (u *UserController) listHandler(c *gin.Context) {
	users, err := u.uc.FindAllUser()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Ok",
		"data":    users,
	})
}
func (u *UserController) getHandler(c *gin.Context)    {}
func (u *UserController) updateHandler(c *gin.Context) {}
func (u *UserController) deleteHandler(c *gin.Context) {}

func (u *UserController) Route() {
	rg := u.router.Group("/api/v1")
	rg.POST("/users", u.authMiddleware.RequireToken("admin"), u.createHandler)
	rg.GET("/users", u.authMiddleware.RequireToken("admin"), u.listHandler)
	rg.GET("/users/:id", u.authMiddleware.RequireToken("admin"), u.getHandler)
	rg.PUT("/users", u.authMiddleware.RequireToken("admin"), u.updateHandler)
	rg.DELETE("/users/:id", u.authMiddleware.RequireToken("admin"), u.deleteHandler)
}

func NewUserController(uc usecase.UserUseCase, r *gin.Engine, am middleware.AuthMiddleware) *UserController {
	controller := &UserController{
		uc:             uc,
		router:         r,
		authMiddleware: am,
	}
	return controller
}
