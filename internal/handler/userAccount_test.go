package handler

import(
	"bitbucket.org/falabellafif/ExampleProject/internal/mock/controller"
	"bytes"
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
	"net/http/httptest"
	"encoding/json"
	"testing"
	. "github.com/smartystreets/goconvey/convey"
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
		var m = make(map[string]string)
		Convey("Test binding error", func() {
			m["usr_name"] = "Rodrigo"
			m["usr_pwd"] = "123123"
			req, _ := http.NewRequest("POST","/register",bytes.NewBufferString(json.Marshal(m)))
			engine.ServeHTTP(res, req)

			So(res.Code, ShouldEqual, 400)
		})
	})
}
