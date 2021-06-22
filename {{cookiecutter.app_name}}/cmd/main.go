package main

import (
	"fmt"

	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/pkg/users/application"
	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/pkg/users/infrastructure/adapters/inmemory"
	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/pkg/users/interfaces/handlers"
)

func main() {
	fmt.Println("Hola papi")

	{% if cookiecutter.use_inmemory_storage == "y" %}
	repo := inmemory.NewInMemoryStorage()
	{% else %}
	{% endif %}

	usecases := application.NewUserUsecase(repo)

	{% if cookiecutter.use_echo_api == "y" %}
	e := handlers.NewEchoServer(usecases)
	e.Logger.Fatal(e.Start(":8080"))
	{% else %}
	{% endif %}
}
