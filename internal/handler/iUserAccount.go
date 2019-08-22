package handler

import "net/http"

type IUserAccount interface {
	RegisterUserAccount() func(w http.ResponseWriter, r *http.Request)
}
