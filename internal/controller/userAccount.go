package controller

import (
	"bitbucket.org/falabellafif/ExampleProject/internal/repository"
	"bitbucket.org/falabellafif/ExampleProject/internal/repository/model"
	"errors"
)

type UserAccount struct {
	repository repository.IUserAccount
}

func (u *UserAccount) RegisterUser(accounts *model.UserAccounts) (*model.UserAccounts, error){

	if userAccountExist, err := u.repository.FindUserAccountByUserName(accounts); err != nil {
		return nil, err
	} else if userAccountExist{
		return nil, errors.New("El usuario ya está registrado")
	}
	if newUserAccount, err := u.repository.CreateUserAccount(accounts); err != nil {
		return nil, err
	} else {
		return newUserAccount, nil
	}
}

func NewUserAccount(u *repository.UserAccount) *UserAccount {
	return &UserAccount{repository: u}
}