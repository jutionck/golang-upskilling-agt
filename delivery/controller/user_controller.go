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

func NewUserController(uc usecase.UserUseCase, r *gin.Engine, am middleware.AuthMiddleware) *UserController {
	controller := &UserController{
		uc:             uc,
		router:         r,
		authMiddleware: am,
	}
	// /api/v1
	rg := r.Group("/api/v1")
	rg.POST("/users", am.RequireToken("admin"), controller.createHandler)
	rg.GET("/users", am.RequireToken("admin"), controller.listHandler)
	rg.GET("/users/:id", am.RequireToken("admin"), controller.getHandler)
	rg.PUT("/users", am.RequireToken("admin"), controller.updateHandler)
	rg.DELETE("/users/:id", am.RequireToken("admin"), controller.deleteHandler)
	return controller
}
