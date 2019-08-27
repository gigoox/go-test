package route

import (
	"bitbucket.org/falabellafif/ExampleProject/internal/handler"
	"github.com/gin-gonic/gin"
)

type UserAccount struct {
	handler handler.IUserAccount
	engine *gin.Engine
}

func (u *UserAccount) SetRoute(){
	route := u.engine.Group("api/v1/")
	{
		route.POST("/register", u.handler.RegisterUserAccount)
	}
}

func NewUserAccount(h *handler.UserAccount, e *gin.Engine) *UserAccount{
	return &UserAccount{
		handler: h,
		engine: e,
	}
}
