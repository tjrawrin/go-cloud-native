package app

import "go-cloud-native/util/logger"

// App ...
type App struct {
	logger *logger.Logger
}

// New ...
func New(logger *logger.Logger) *App {
	return &App{logger: logger}
}

// Logger ...
func (app *App) Logger() *logger.Logger {
	return app.logger
}
