package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/pkg/users/application"
	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/pkg/users/interfaces"
)

// APIHandler ...
type APIHandler struct {
	Usecases application.Usecases
}

// ErrCreatingUser ...
var ErrCreatingUser = errors.New("Error creating new User")

// ErrLoginUser ...
var ErrLoginUser = errors.New("Error login User")

// ErrLogoutUser ...
var ErrLogoutUser = errors.New("Error logout User")

// ErrGettingUser ...
var ErrGettingUser = errors.New("Error getting the user")

// ErrDeletingUser ...
var ErrDeletingUser = errors.New("Error deleting the user")

// CustomValidator struct
type CustomValidator struct {
	validator *validator.Validate
}

// Validate method
func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return nil
}

// NewEchoServer creates echo server
func NewEchoServer(usecases application.Usecases) *echo.Echo {
	handler := &APIHandler{
		Usecases: usecases,
	}
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello ECHO followers")
	})

	e.POST("/users", handler.RegisterUser)
	e.POST("/users/login", handler.LoginUser)
	e.GET("/users/logout", handler.LogoutUser)
	e.GET("/users/:id", handler.GetUser)
	e.DELETE("/users/:id", handler.DeleteUser)

	return e
}

// RegisterUser echo handler
func (h *APIHandler) RegisterUser(c echo.Context) error {
	req := new(interfaces.RegisterUserRequest)
	if err := c.Bind(req); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return ErrCreatingUser
	}
	if err := c.Validate(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user, err := h.Usecases.RegisterUser(req.Name, req.Surname, req.Email, req.Password, req.Age)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return ErrCreatingUser
	}

	userData, err := json.Marshal(user)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return ErrGettingUser
	}

	return c.String(http.StatusOK, fmt.Sprintf("User: %v", string(userData)))
}

// LoginUser echo handler
func (h *APIHandler) LoginUser(c echo.Context) error {
	req := new(interfaces.LoginUserRequest)
	if err := c.Bind(req); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return ErrLoginUser
	}
	if err := c.Validate(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user, err := h.Usecases.LoginUser(req.Email, req.Password)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return ErrLoginUser
	}

	cookie := new(http.Cookie)
	cookie.Name = "user"
	cookie.Value = user.ID
	cookie.Expires = time.Now().Add(24 * time.Hour)
	c.SetCookie(cookie)

	return c.String(http.StatusOK, fmt.Sprintf("User successfully login"))
}

// LogoutUser echo handler
func (h *APIHandler) LogoutUser(c echo.Context) error {
	userCookie, err := c.Cookie("user")
	if err != nil {
		return err
	}

	err = h.Usecases.LogoutUser(userCookie.Value)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return ErrLogoutUser
	}

	// Setting MaxAge<0 means delete cookie now
	cookie := new(http.Cookie)
	cookie.Name = "user"
	cookie.MaxAge = -1
	c.SetCookie(cookie)

	return c.String(http.StatusOK, fmt.Sprintf("User successfully logout"))
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

	return c.String(http.StatusOK, fmt.Sprintf("User: %v", string(userData)))
}

// DeleteUser echo handler
func (h *APIHandler) DeleteUser(c echo.Context) error {
	if err := h.Usecases.DeleteUser(c.Param("id")); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return ErrDeletingUser
	}

	return c.String(http.StatusOK, fmt.Sprintf("User %v Deleted", c.Param("id")))
}
