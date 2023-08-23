package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jutionck/golang-upskilling-agt/model"
	"github.com/jutionck/golang-upskilling-agt/usecase"
)

type UserController struct {
	uc     usecase.UserUseCase
	router *gin.Engine
}

func (u *UserController) createHandler(c *gin.Context) {
	var payload model.User
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	payload.Id = uuid.New().String()
	if err := u.uc.RegisterNewUser(payload); err != nil {
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

func NewUserController(uc usecase.UserUseCase, r *gin.Engine) *UserController {
	controller := &UserController{
		uc:     uc,
		router: r,
	}
	// /api/v1
	rg := r.Group("/api/v1")
	rg.POST("/users", controller.createHandler)
	rg.GET("/users", controller.listHandler)
	rg.GET("/users/:id", controller.getHandler)
	rg.PUT("/users", controller.updateHandler)
	rg.DELETE("/users/:id", controller.deleteHandler)
	return controller
}
