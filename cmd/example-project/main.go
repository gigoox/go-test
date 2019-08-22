package main

import (
	"bitbucket.org/falabellafif/ExampleProject/internal/config"
	"bitbucket.org/falabellafif/ExampleProject/internal/controller"
	"bitbucket.org/falabellafif/ExampleProject/internal/handler"
	"bitbucket.org/falabellafif/ExampleProject/internal/repository"
	"github.com/go-chi/chi"
	"go.uber.org/dig"
	"gopkg.in/go-playground/validator.v9"
)

func main()  {
	container:= createContainer()
	if err := container.Invoke(func(server *Server){
		server.StartServer()
	}); err != nil {
		panic(err)
	}
}

func createContainer() (*dig.Container) {
	container := dig.New()
	if err := container.Provide(config.NewDb); err != nil {
		panic(err)
	}
	if err := container.Provide(CreateDatabaseConnection); err != nil {
		panic(err)
	}
	if err := container.Provide(repository.NewUserAccount); err != nil {
		panic(err)
	}
	if err := container.Provide(controller.NewUserAccount); err != nil {
		panic(err)
	}
	if err := container.Provide(validator.New); err != nil {
		panic(err)
	}
	if err := container.Provide(handler.NewUserAccount); err != nil {
		panic(err)
	}
	if err := container.Provide(NewServer); err != nil {
		panic(err)
	}
	if err := container.Provide(chi.NewRouter); err != nil {
		panic(err)
	}
	return container
}


