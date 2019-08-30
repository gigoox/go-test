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
		engine := gin.Default()
		userAccount := &UserAccount{
			controller:controllerMock,
			validator: validator.New(),
		}
		engine.POST("/register", userAccount.RegisterUserAccount)
		res := httptest.NewRecorder()
		Convey("Test binding error", func() {
			mcPostBody := map[string]interface{}{
				"user_passwd": "Is this a test post for MutliQuestion?",
			}
			body, _ := json.Marshal(mcPostBody)
			req, _ := http.NewRequest("POST","/register", strings.NewReader(string(body)))
			req.Header.Add("Content-Type", "application/json")
			engine.ServeHTTP(res, req)

			So(res.Code, ShouldEqual, 400)
		})
	})
}
