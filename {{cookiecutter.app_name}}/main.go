package main

import (
	"fmt"

	"github.com/rubenesp87/cookiecutter-go-hexagonal-architecture/pkg/users/application"
	"github.com/rubenesp87/cookiecutter-go-hexagonal-architecture/pkg/users/infraestructure/adapters/inmemory"
	"github.com/rubenesp87/cookiecutter-go-hexagonal-architecture/pkg/users/interfaces/handlers"
)

func main() {
	// {% if cookiecutter.use_cobra_cmd == "y" %}
	fmt.Println("Hola papi")
	// {% else %}
	// fmt.Println("Hola mami")
	// {% endif %}

	repo := inmemory.NewInMemoryStorage()

	usecases := application.NewUserUsecase(repo)

	e := handlers.NewEchoServer(usecases)
	e.Logger.Fatal(e.Start(":8080"))
}
