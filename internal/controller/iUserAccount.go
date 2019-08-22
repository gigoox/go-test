package controller

import "bitbucket.org/falabellafif/ExampleProject/internal/repository/model"

type IUserAccount interface {
	RegisterUser(accounts *model.UserAccounts) (*model.UserAccounts, error)
}
