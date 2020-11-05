package repositories

import (
	"github.com/olucvolkan/todoApp/api/models"
	"github.com/olucvolkan/todoApp/app"
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
