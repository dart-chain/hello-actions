package app

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/dart-chain/hello-actions/internal/config"
	"github.com/dart-chain/hello-actions/internal/handlers"
	"github.com/dart-chain/hello-actions/internal/httpx"
	"gitlab.com/tim_de/dartlog/v2/dartlog"
)

type App struct {
	Name    string
	Version string

	Server *httpx.Server
	Router *httpx.Router
	Logger *dartlog.Logger
}

func New(name, version string, config *config.AppCfg, logger *dartlog.Logger) *App {
	router := httpx.NewRouter(config.IsDebug, logger)

	server := httpx.NewServer(config.Host, config.Port, router)

	return &App{
		Name:    name,
		Version: version,

		Server: server,
		Router: router,
		Logger: logger,
	}
}

func (a *App) setRoutes() {
	healthAPI := handlers.NewHealthAPI()
	healthAPI.RegisterRoutes(a.Router.Group("/"), a.Name)
}

func (a *App) Run() error {
	a.setRoutes()

	errChan := make(chan error, 1)
	go func() {
		a.Logger.Log(dartlog.INFO, "App", "Starting HTTP server", map[string]any{
			"name":    a.Name,
			"version": a.Version,
			"address": a.Server.Address(),
		})
		errChan <- a.Server.Run()
	}()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

	select {
	case err := <-errChan:
		a.Logger.Log(dartlog.ERROR, "App", "Runtime error occurred", map[string]any{"error": err.Error()})
		return err
	case <-signalChan:
		return a.Server.Shutdown()
	}
}
