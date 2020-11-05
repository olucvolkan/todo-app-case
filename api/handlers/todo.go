package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/olucvolkan/todoApp/api/models"
	"github.com/olucvolkan/todoApp/api/schema"
)

// GetTodos godoc
// @Summary List all todos
// @Description List todo list
// @Tags todos
// @Accept  json
// @Produce  json
// @Router /todo-list [get]
func (h *Handlers) GetTodos(c *gin.Context) {
	var todos []models.Todo
	todos, err := h.todoRepo.GetAll()

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, presentTodosWithSchema(todos))
	}
}

// CreateTodoRequest stores request payload for creating todo request
type CreateTodoRequest struct {
	Description string `json:"description" binding:"required"`
	Status      string `json:"status" binding:"required"`
}

// CreateTodo godoc
// @Summary Create a todo
// @Description  Create a todo
// @Tags todos
// @Accept  json
// @Produce  json
// @Param tasks body models.Todo  true "todo info"
// @Router /create-todo [post]
func (h *Handlers) CreateTodo(c *gin.Context) {
	var req CreateTodoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	todo := models.Todo{
		Description: req.Description,
		Status:      req.Status,
	}

	if err := h.DB.Create(&todo).Error; err != nil {
		log.Fatalf("can't create todo, err: %+v", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusCreated, todo)
}

// DeleteTodoRequest stores payload for delete request
type DeleteTodoRequest struct {
	ID uint `json:"id" binding:"required"`
}

// DeleteTodo godoc
// @Summary Delete  a todo
// @Description  Delete  a todo
// @Tags todos
// @Accept  json
// @Produce  json
// @Param tasks body models.DeleteTodoRequest  true "true"
// @Router /delete-todo [post]
func (h *Handlers) DeleteTodo(c *gin.Context) {
	var req DeleteTodoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if err := h.DB.Delete(&models.Todo{}, req.ID).Error; err != nil {
		log.Fatalf("can't delete todo, err: %+v", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.String(http.StatusNoContent, "")
}

// UpdateTodoRequest stores payload for update request
type UpdateTodoRequest struct {
	ID     uint   `json:"id" binding:"required"`
	Status string `json:"status" binding:"required"`
}

// UpdateTodo godoc
// @Summary Update   a todo
// @Description  Update  a todo
// @Tags todos
// @Accept  json
// @Produce  json
// @Param tasks body models.UpdateTodoRequest true "todo info"
// @Router /update-todo [post]
func (h *Handlers) UpdateTodo(c *gin.Context) {
	var req UpdateTodoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Println(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	var todo *models.Todo
	if err := h.DB.Model(&models.Todo{}).Where("id = ?", req.ID).Update("status", req.Status).Error; err != nil {
		log.Fatalf("can't update todo, err: %+v", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, todo)
}

func presentTodosWithSchema(todos []models.Todo) schema.TodoListSchema {
	var listSchema = schema.TodoListSchema{
		Todo:       schema.Todo{Title: "todo", Items: make([]schema.Item, 0)},
		Inprogress: schema.InProgress{Title: "in-progress", Items: make([]schema.Item, 0)},
		Done:       schema.Done{Title: "done", Items: make([]schema.Item, 0)},
	}
	for _, todo := range todos {
		addToList(&listSchema, todo.ID, todo.Description, todo.Status)
	}

	return listSchema
}

func addToList(todoListSchema *schema.TodoListSchema, todoID uint, todoText string, todoStatus string) {
	var todoIDString = strconv.FormatUint(uint64(todoID), 10)
	item := schema.Item{ID: todoIDString, Description: todoText, Status: todoStatus}

	switch todoStatus {
	case "todo":
		todoListSchema.Todo.Items = append(todoListSchema.Todo.Items, item)
	case "in-progress":
		todoListSchema.Inprogress.Items = append(todoListSchema.Inprogress.Items, item)
	case "done":
		todoListSchema.Done.Items = append(todoListSchema.Done.Items, item)
	}
}
