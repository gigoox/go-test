package main

import (
	"bitbucket.org/falabellafif/ExampleProject/internal/route"
	"fmt"
	"github.com/gin-gonic/gin"
)

type Server struct {
	engine *gin.Engine
	route route.IUserAccount
}

func (s *Server) StartServer(){
	s.route.SetRoute()
	if err := s.engine.Run(":9090"); err != nil {
		panic(err)
	}
	fmt.Print("Servidor levantado en el puerto 9090")
}

func NewServer(e *gin.Engine, r *route.UserAccount) *Server{
	return &Server{
		engine:e,
		route:r,
	}
}