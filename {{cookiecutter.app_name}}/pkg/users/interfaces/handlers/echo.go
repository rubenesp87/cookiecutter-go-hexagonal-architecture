package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/rubenesp87/cookiecutter-go-hexagonal-architecture/pkg/users/application"
	"github.com/rubenesp87/cookiecutter-go-hexagonal-architecture/pkg/users/interfaces"
)

// APIHandler ...
type APIHandler struct {
	Usecases application.Usecases
}

// ErrCreatingUser ...
var ErrCreatingUser = errors.New("Error creating new User")

// ErrGettingUser ...
var ErrGettingUser = errors.New("Error getting the user")

// ErrDeletingUser ...
var ErrDeletingUser = errors.New("Error deleting the user")

// NewEchoServer creates echo server
func NewEchoServer(usecases application.Usecases) *echo.Echo {
	handler := &APIHandler{
		Usecases: usecases,
	}
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello ECHO followers")
	})

	e.POST("/users", handler.CreateUser)
	e.GET("/users/:id", handler.GetUser)
	e.DELETE("/users/:id", handler.DeleteUser)

	return e
}

// CreateUser echo handler
func (h *APIHandler) CreateUser(c echo.Context) error {
	req := new(interfaces.CreateUserRequest)
	if err := c.Bind(req); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return ErrCreatingUser
	}

	user, err := h.Usecases.CreateUser(req.Name, req.Surname, req.Email, req.Age)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return ErrCreatingUser
	}

	userData, err := json.Marshal(user)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return ErrGettingUser
	}

	return c.String(http.StatusOK, string(userData))
}

// GetUser echo handler
func (h *APIHandler) GetUser(c echo.Context) error {
	user, err := h.Usecases.GetUser(c.Param("id"))
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return ErrGettingUser
	}

	userData, err := json.Marshal(user)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return ErrGettingUser
	}

	return c.String(http.StatusOK, string(userData))
}

// DeleteUser echo handler
func (h *APIHandler) DeleteUser(c echo.Context) error {
	if err := h.Usecases.DeleteUser(c.Param("id")); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return ErrDeletingUser
	}

	return c.String(http.StatusOK, fmt.Sprintf("User %v Deleted", c.Param("id")))
}
