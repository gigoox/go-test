package repository

import(
	"bitbucket.org/falabellafif/ExampleProject/internal/repository/model"
)

type IUserAccount interface {
	//FindByUserName(userName string) (*model.UserAccounts, error)
	CreateUserAccount(account *model.UserAccounts) (*model.UserAccounts, error)
	FindUserAccountByUserName(account *model.UserAccounts) (bool, error)
	//UpdateUserPassword(account model.UserAccounts) (*model.UserAccounts, error)
	//UpdateUserLogin()
}
