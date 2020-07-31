package app

import (
	"github.com/jinzhu/gorm"

	"go-cloud-native/util/logger"
)

// App ...
type App struct {
	logger *logger.Logger
	db     *gorm.DB
}

// New ...
func New(logger *logger.Logger, db *gorm.DB) *App {
	return &App{
		logger: logger,
		db:     db,
	}
}

// Logger ...
func (app *App) Logger() *logger.Logger {
	return app.logger
}
