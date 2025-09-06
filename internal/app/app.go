package app

import "gitlab.com/tim_de/dartlog/v2/dartlog"

type App struct {
	Name    string
	Version string

	Logger *dartlog.Logger
}

func New(name, version string, logger *dartlog.Logger) *App {
	return &App{
		Name:    name,
		Version: version,
		Logger:  logger,
	}
}

func (a *App) Run() error {
	a.Logger.Log(dartlog.INFO, "App", "Application is running", map[string]any{
		"name":    a.Name,
		"version": a.Version,
	})
	// Application logic goes here
	return nil
}
