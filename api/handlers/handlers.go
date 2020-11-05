package handlers

import (
	"github.com/olucvolkan/todo-app-case/api/repositories"
	"github.com/olucvolkan/todo-app-case/app"
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
