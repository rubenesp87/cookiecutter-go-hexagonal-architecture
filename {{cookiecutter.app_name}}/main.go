package main

import (
	"fmt"

	"github.com/rubenesp87/cookiecutter-go-hexagonal-architecture/pkg/users/application"
	"github.com/rubenesp87/cookiecutter-go-hexagonal-architecture/pkg/users/infraestructure/adapters/inmemory"
	"github.com/rubenesp87/cookiecutter-go-hexagonal-architecture/pkg/users/interfaces/handlers"
)

func main() {
	fmt.Println("Hola papi")

	// {% if cookiecutter.use_inmemory_storage == "y" %}
	repo := inmemory.NewInMemoryStorage()
	// {% else %}
	// {% endif %}

	usecases := application.NewUserUsecase(repo)

	// {% if cookiecutter.use_echo_api == "y" %}
	e := handlers.NewEchoServer(usecases)
	e.Logger.Fatal(e.Start(":8080"))
	// {% else %}
	// {% endif %}
}
