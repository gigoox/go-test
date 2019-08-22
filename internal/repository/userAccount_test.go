package repository

import (
	"bitbucket.org/falabellafif/ExampleProject/internal/mock/repository"
	"bitbucket.org/falabellafif/ExampleProject/internal/repository/model"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSpec_CreateUserAccount(t *testing.T) {
	Convey("Test insert", t, func() {
		Convey("Test insert failed", func() {
			engine := &repository.EngineMock{
				Err:true,
			}
			userAccount := UserAccount{engine}
			_, err := userAccount.CreateUserAccount(&model.UserAccounts{UserName: "Test", UserPasswd: "uno23"})
			So(err, ShouldNotBeNil)
		})
		Convey("Test insert ok", func() {
			engine := &repository.EngineMock{
				Err:false,
			}
			userAccount := UserAccount{engine}
			newUserAccount, _ := userAccount.CreateUserAccount(&model.UserAccounts{UserName: "Test", UserPasswd: "uno23"})
			So(newUserAccount, ShouldNotBeNil)
			So(newUserAccount.Id, ShouldEqual, 1)
		})
	})
}

func TestSpec_findUserAccountByUserName(t *testing.T) {
	Convey("Test find by user name", t, func() {
		Convey("Test find by user name failed", func() {
			engine := &repository.EngineMock{
				Err:true,
			}
			userAccount := UserAccount{engine}
			_, err := userAccount.FindUserAccountByUserName(&model.UserAccounts{UserName: "Test", UserPasswd: "uno23"})
			So(err, ShouldNotBeNil)
		})
		Convey("Test find by user name ok user account found", func() {
			engine := &repository.EngineMock{
				Err:false,
				UserExist: true,
			}
			userAccount := UserAccount{engine}
			isAnyUserAccount, _ := userAccount.FindUserAccountByUserName(&model.UserAccounts{UserName: "Test", UserPasswd: "uno23"})
			So(isAnyUserAccount, ShouldEqual, true)
		})
		Convey("Test find by user name ok user account not found", func() {
			engine := &repository.EngineMock{
				Err:false,
				UserExist: false,
			}
			userAccount := UserAccount{engine}
			isAnyUserAccount, _ := userAccount.FindUserAccountByUserName(&model.UserAccounts{UserName: "Test", UserPasswd: "uno23"})
			So(isAnyUserAccount, ShouldEqual, false)
		})
	})
}