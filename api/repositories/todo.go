package repositories

import (
	"github.com/olucvolkan/todo-app-case/api/models"
	"github.com/olucvolkan/todo-app-case/app"
)

// TodoRepo holds methods to interact with todo records
type TodoRepo struct {
	*app.App
}

// GetAll returns list of all todos
func (tr *TodoRepo) GetAll() ([]models.Todo, error) {
	var todos []models.Todo

	if err := tr.DB.Find(&todos).Error; err != nil {
		return nil, err
	}
	return todos, nil
}
