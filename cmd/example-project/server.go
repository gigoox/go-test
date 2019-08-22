package main

import (
	"bitbucket.org/falabellafif/ExampleProject/internal/handler"
	"github.com/go-chi/chi"
	"net/http"
)

type Server struct {
	router *chi.Mux
	handler handler.IUserAccount
}

func (s *Server) StartServer(){
	s.router.Post("/test", s.handler.RegisterUserAccount())
	http.ListenAndServe(":3000", s.router)
}

func NewServer(r *chi.Mux, h *handler.UserAccount) *Server{
	return &Server{
		router:r,
		handler: h,
	}
}