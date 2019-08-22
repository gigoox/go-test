package controller

import (
	"bitbucket.org/falabellafif/ExampleProject/internal/mock/repository"
	"bitbucket.org/falabellafif/ExampleProject/internal/repository/model"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSpec_RegisterUser(t *testing.T) {
	Convey("Test register account", t, func() {
		Convey("Test register finding account error", func() {
			engine := &repository.UserAccountMock{
				ErrorFindAccount: true,
			}
			userAccount := UserAccount{engine}
			_, err := userAccount.RegisterUser(&model.UserAccounts{UserName: "Test", UserPasswd: "uno23"})
			So(err, ShouldNotBeNil)
		})
		Convey("Test register finding account exist", func() {
			engine := &repository.UserAccountMock{
				ErrorFindAccount: false,
				UserExist: true,
			}
			userAccount := UserAccount{engine}
			_, err := userAccount.RegisterUser(&model.UserAccounts{UserName: "Test", UserPasswd: "uno23"})
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldEqual, "El usuario ya está registrado" )
		})
		Convey("Test register creating account error", func() {
			engine := &repository.UserAccountMock{
				ErrorFindAccount: false,
				UserExist: false,
				ErrorCreatAccount: true,
			}
			userAccount := UserAccount{engine}
			_, err := userAccount.RegisterUser(&model.UserAccounts{UserName: "Test", UserPasswd: "uno23"})
			So(err, ShouldNotBeNil)
		})
		Convey("Test register creating account ok", func() {
			engine := &repository.UserAccountMock{
				ErrorFindAccount: false,
				UserExist: false,
				ErrorCreatAccount: false,
			}
			userAccount := UserAccount{engine}
			accountRegistered, err := userAccount.RegisterUser(&model.UserAccounts{UserName: "Test", UserPasswd: "uno23"})
			So(err, ShouldBeNil)
			So(accountRegistered, ShouldNotBeNil)
		})
	})
}
