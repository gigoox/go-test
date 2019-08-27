package handler

import (
	"github.com/gin-gonic/gin"
)

type IUserAccount interface {
	RegisterUserAccount(c *gin.Context)
}
