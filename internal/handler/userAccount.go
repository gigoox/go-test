package handler

import (
	"bitbucket.org/falabellafif/ExampleProject/internal/controller"
	"bitbucket.org/falabellafif/ExampleProject/internal/repository/model"
	"encoding/json"
	"fmt"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
)

type UserAccount struct {
	controller controller.IUserAccount
	validator *validator.Validate
}

func (u *UserAccount) RegisterUserAccount() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var userAccount = new(model.UserAccounts)
		if err := json.NewDecoder(r.Body).Decode(userAccount); err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, `{"state":"Error","message":%q}`, err.Error())
			return
		}
		if err := u.validator.Struct(userAccount); err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, `{"state":"Error","message":%q}`, err.Error())
			return
		}
		userAccount, err := u.controller.RegisterUser(userAccount)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, `{"state":"Error","message":%q}`, err.Error())
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		body,_ := json.Marshal(userAccount)
		fmt.Fprintf(w, `{"state":"Ok","message": "Account created", "Account":%v}`, string(body))
	}
}

func NewUserAccount(u *controller.UserAccount, v *validator.Validate) *UserAccount{
	return &UserAccount{
		controller: u,
		validator: v,
	}
}
