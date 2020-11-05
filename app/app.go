package app

import (
	"github.com/jinzhu/gorm"
)

// App is container for app dependencies
type App struct {
	DB *gorm.DB
}

// New returns new instance of App
func New(DB *gorm.DB) *App {
	return &App{DB: DB}
}
