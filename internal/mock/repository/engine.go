package repository

import "errors"

type EngineMock struct {
	Err bool
	UserExist bool
}

func (e *EngineMock) Insert(beans ...interface{}) (int64, error)  {
	if e.Err {
		return 0, errors.New("Error")
	}
	return 1, nil
}

func (e *EngineMock) Exist(bean ...interface{}) (bool, error){
	if e.Err {
		return false, errors.New("Error")
	}
	if e.UserExist {
		return true, nil
	}
	return false, nil
}
