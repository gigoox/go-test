package repository

import (
	"bitbucket.org/falabellafif/ExampleProject/internal/repository/model"
	"github.com/go-xorm/xorm"
)

type UserAccount struct {
	engine IEngine
}

func (u *UserAccount) CreateUserAccount(account *model.UserAccounts) (*model.UserAccounts, error){
	userAccountId, err := u.engine.Insert(account)
	if err != nil {
		return nil, err
	}
	account.Id = userAccountId
	return account, nil
}

func (u *UserAccount) FindUserAccountByUserName(account *model.UserAccounts) (bool, error) {
	isAnyUserAccount, err := u.engine.Exist(account)
	if err != nil {
		return false, err
	}
	return isAnyUserAccount, nil
}

func NewUserAccount(e *xorm.Engine) *UserAccount{
	return &UserAccount{engine: e}
}
