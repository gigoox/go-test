package controller

import (
	"bitbucket.org/falabellafif/ExampleProject/internal/repository/model"
	"errors"
)

type ControllerMock struct {
	ErrorFindAccount bool
	UserExist bool
	ErrorCreatAccount bool
}

func (u *ControllerMock) RegisterUser(accounts *model.UserAccounts) (*model.UserAccounts, error){
	if u.ErrorFindAccount {
		return nil, errors.New("")
	}
	if u.UserExist {
		return nil, errors.New("El usuario ya está registrado")
	}
	if u.ErrorCreatAccount {
		return nil, errors.New("")
	}
	return new(model.UserAccounts), nil
}
