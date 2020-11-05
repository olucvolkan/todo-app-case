package server

import (
	"github.com/gin-gonic/gin"
	"github.com/olucvolkan/todo-app-case/api/handlers"
)

// New creates new instance of Gin and configures it
func New(h *handlers.Handlers) *gin.Engine {
	router := gin.Default()
	router.Use(corsMiddleware())
	router.GET("/todo-list", h.GetTodos)
	router.POST("/create-todo", h.CreateTodo)
	router.POST("/delete-todo", h.DeleteTodo)
	router.POST("/update-todo", h.UpdateTodo)

	return router
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST,HEAD,PATCH,OPTIONS,GET,PUT,DELETE")
		c.Next()
	}
}
