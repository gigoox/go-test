package model

import (
	"time"
)

type UserAccounts struct {
	Id   int64		   		`xorm:"primary_key autoincr 'id'" json:"id"`
	UserName string    		`xorm:"varchar(25) not null unique 'usr_name'" json:"user_name" validate:"required,min=5,max=25"`
	UserPasswd string  		`xorm:"varchar(100) not null 'usr_pwd'" json:"user_passwd" validate:"required"`
	UserCreated time.Time	`xorm:"created not null 'created'" json:"user_created"`
	LastLogin time.Time 	`xorm:"timestamp 'last_login'" json:"last_login"`
}

