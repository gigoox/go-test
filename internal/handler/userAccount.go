package handler

import (
	"bitbucket.org/falabellafif/ExampleProject/internal/controller"
	"bitbucket.org/falabellafif/ExampleProject/internal/repository/model"
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
)

type UserAccount struct {
	controller controller.IUserAccount
	validator *validator.Validate
}

func (u *UserAccount) RegisterUserAccount(c *gin.Context) {
	var userAccount = new(model.UserAccounts)
	if err := c.Bind(&userAccount); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{ "message": err.Error()})
		return
	}
	if err := u.validator.Struct(userAccount); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{ "message": err.Error()})
		return
	}
	userAccount, err := u.controller.RegisterUser(userAccount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{ "message": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, &userAccount)
}

func NewUserAccount(u *controller.UserAccount, v *validator.Validate) *UserAccount{
	return &UserAccount{
		controller: u,
		validator: v,
	}
}
