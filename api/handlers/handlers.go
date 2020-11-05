package handlers

import (
	"github.com/olucvolkan/todoApp/api/repositories"
	"github.com/olucvolkan/todoApp/app"
)

// Handlers holds all controller methods
type Handlers struct {
	*app.App
	todoRepo *repositories.TodoRepo
}

// New returns new instance of controller
func New(a *app.App, todoRepo *repositories.TodoRepo) *Handlers {
	return &Handlers{App: a, todoRepo: todoRepo}
}
