package test

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/olucvolkan/todoApp/api/config"
	"github.com/olucvolkan/todoApp/api/handlers"
	"github.com/olucvolkan/todoApp/api/models"
	"github.com/olucvolkan/todoApp/api/repositories"
	"github.com/olucvolkan/todoApp/api/server"
	"github.com/olucvolkan/todoApp/app"
	"github.com/stretchr/testify/assert"
)

var err error

func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func createTestEnvironment() *gin.Engine {
	config := config.New()

	ensureDBExists(config)

	gormDB, err := gorm.Open("mysql", config.DBUrl())
	if err != nil {
		fmt.Println(fmt.Errorf("Can't connect to database, err: %v", err))
		return nil
	}
	gormDB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(models.Todo{})

	if err != nil {
		fmt.Println("status: ", err)
	}
	gormDB.AutoMigrate(&models.Todo{})

	app := app.New(gormDB)
	todoRepo := &repositories.TodoRepo{App: app}
	handlers := handlers.New(app, todoRepo)
	server := server.New(handlers)
	return server
}

func ensureDBExists(config *config.Config) {
	db, err := sql.Open("mysql", config.DBUrlWithoutDBName())

	if err != nil {
		fmt.Println(err)
		fmt.Println("can't connect database for creating table")
		return
	}
	defer db.Close()

	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS " + config.DBName + ";")

	if err != nil {
		fmt.Println(err)
		fmt.Println(err.Error())
	} else {
		fmt.Println("Successfully created database or updated")
	}
}

func TestTodoCRUD(t *testing.T) {
	router := createTestEnvironment()

	t.Run("Create Todo endpoint", func(t *testing.T) {
		todo, _ := json.Marshal(handlers.CreateTodoRequest{Description: "Test Record", Status: "Todo"})

		fmt.Println(string(todo))
		req, err := http.NewRequest("POST", "/create-todo", bytes.NewReader(todo))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, nil, err)
		assert.Equal(t, http.StatusCreated, w.Code)
	})

	t.Run("List Todo endpoint", func(t *testing.T) {
		w := performRequest(router, "GET", "/todo-list")
		assert.Equal(t, nil, err)
		assert.Equal(t, http.StatusOK, w.Code)

	})

	t.Run("Update Todo endpoint", func(t *testing.T) {
		todo, _ := json.Marshal(handlers.UpdateTodoRequest{ID: 1, Status: "Todo"})

		fmt.Println(string(todo))
		req, err := http.NewRequest("POST", "/update-todo", bytes.NewReader(todo))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, nil, err)
		assert.Equal(t, http.StatusOK, w.Code)
	})
	t.Run("Delete Todo endpoint", func(t *testing.T) {

		todo, _ := json.Marshal(handlers.DeleteTodoRequest{ID: 1})

		fmt.Println(string(todo))
		req, err := http.NewRequest("POST", "/delete-todo", bytes.NewReader(todo))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, nil, err)
		assert.Equal(t, http.StatusNoContent, w.Code)
	})
}
