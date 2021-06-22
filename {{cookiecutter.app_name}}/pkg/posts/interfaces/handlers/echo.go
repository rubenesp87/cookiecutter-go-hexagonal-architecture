package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/pkg/posts/application"
	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/pkg/posts/interfaces"
)

// APIHandler ...
type APIHandler struct {
	Usecases application.Usecases
}

// ErrCreatingPost ...
var ErrCreatingPost = errors.New("Error creating new post")

// ErrGettingPost ...
var ErrGettingPost = errors.New("Error getting the post")

// ErrDeletingPost ...
var ErrDeletingPost = errors.New("Error deleting the post")

// NewEchoServer creates echo server
func NewEchoServer(usecases application.Usecases) *echo.Echo {
	handler := &APIHandler{
		Usecases: usecases,
	}
	e := echo.New()

	e.POST("/posts/:user_id", handler.CreatePost)
	e.GET("/posts/:post_id", handler.GetPost)
	e.DELETE("/posts/:post_id/:user_id", handler.DeletePost)

	return e
}

// CreatePost echo handler
func (h *APIHandler) CreatePost(c echo.Context) error {
	req := new(interfaces.CreatePostRequest)
	if err := c.Bind(req); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return ErrCreatingPost
	}

	post, err := h.Usecases.CreatePost(req.Content, c.Param("user_id"))
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return ErrCreatingPost
	}

	postData, err := json.Marshal(post)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return ErrGettingPost
	}

	return c.String(http.StatusOK, fmt.Sprintf("Post: %v", string(postData)))
}

// GetPost echo handler
func (h *APIHandler) GetPost(c echo.Context) error {
	post, err := h.Usecases.GetPost(c.Param("post_id"))
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return ErrGettingPost
	}

	postData, err := json.Marshal(post)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return ErrGettingPost
	}

	return c.String(http.StatusOK, fmt.Sprintf("Post: %v", string(postData)))
}

// DeletePost echo handler
func (h *APIHandler) DeletePost(c echo.Context) error {
	if err := h.Usecases.DeletePost(c.Param("post_id")); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return ErrDeletingPost
	}

	return c.String(http.StatusOK, fmt.Sprintf("Post %v Deleted", c.Param("post_id")))
}
