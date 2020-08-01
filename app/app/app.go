package app

import (
	"github.com/jinzhu/gorm"
	"gopkg.in/go-playground/validator.v9"

	"go-cloud-native/util/logger"
)

const (
	appErrDataAccessFailure      = "data access failure"
	appErrJSONCreationFailure    = "json creation failure"
	appErrDataCreationFailure    = "data creation failure"
	appErrFormDecodingFailure    = "form decoding failure"
	appErrDataUpdateFailure      = "data update failure"
	appErrFormErrResponseFailure = "form error response failure"
)

// App ...
type App struct {
	logger    *logger.Logger
	db        *gorm.DB
	validator *validator.Validate
}

// New ...
func New(logger *logger.Logger, db *gorm.DB, validator *validator.Validate) *App {
	return &App{
		logger:    logger,
		db:        db,
		validator: validator,
	}
}

// Logger ...
func (app *App) Logger() *logger.Logger {
	return app.logger
}
