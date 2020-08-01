package app

import (
	"github.com/jinzhu/gorm"

	"go-cloud-native/util/logger"
)

const (
	appErrDataAccessFailure   = "data access failure"
	appErrJSONCreationFailure = "json creation failure"
	appErrDataCreationFailure = "data creation failure"
	appErrFormDecodingFailure = "form decoding failure"
	appErrDataUpdateFailure   = "data update failure"
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
