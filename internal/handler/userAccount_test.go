package handler

import (
	"bitbucket.org/falabellafif/ExampleProject/internal/mock/controller"
	"encoding/json"
	"github.com/gin-gonic/gin"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestSpec_RegisterUserAccount(t *testing.T) {
	Convey("Test Register User Account", t, func() {
		controllerMock := &controller.ControllerMock{}
		gin.SetMode(gin.TestMode)
		res := httptest.NewRecorder()
		c, engine := gin.CreateTestContext(res)
		userAccount := &UserAccount{
			controller:controllerMock,
			validator: validator.New(),
		}
		engine.Use(func(c *gin.Context) {
			c.Set("Test", 12)
		})
		engine.POST("/register", userAccount.RegisterUserAccount)
		Convey("Test binding error", func() {
			mcPostBody := map[string]interface{}{
				"user_passwd": "Is this a test post for MutliQuestion?",
			}
			body, _ := json.Marshal(mcPostBody)
			c.Request, _ = http.NewRequest("POST","/register", strings.NewReader(string(body)))
			c.Request.Header.Add("Content-Type", "application/json")
			engine.ServeHTTP(res, c.Request)

			So(res.Code, ShouldEqual, 400)
		})
	})
}
