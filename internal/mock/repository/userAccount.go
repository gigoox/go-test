package repository

import (
	"bitbucket.org/falabellafif/ExampleProject/internal/repository/model"
	"errors"
)

type UserAccountMock struct {
	ErrorFindAccount bool
	UserExist bool
	ErrorCreatAccount bool
}

func (u *UserAccountMock) CreateUserAccount(account *model.UserAccounts) (*model.UserAccounts, error){
	if u.ErrorCreatAccount {
		return nil, errors.New("Error al crear cuenta")
	}
	return new(model.UserAccounts), nil
}

func (u *UserAccountMock) FindUserAccountByUserName(account *model.UserAccounts) (bool, error){
	if u.ErrorFindAccount {
		return true, errors.New("Error al buscar cuenta")
	}
	return u.UserExist, nil
}

