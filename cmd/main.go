package main

import (
	"os"

	"github.com/akamensky/argparse"
	"github.com/dart-chain/hello-actions/internal/app"
	"github.com/dart-chain/hello-actions/internal/config"
	"github.com/dart-chain/hello-actions/internal/utils"
	"gitlab.com/tim_de/dartlog/v2/dartlog"
)

var (
	version string
)

func main() {
	logger := dartlog.New("hello-actions", true, false)

	if !utils.IsVersionValid(version) {
		logger.Log(dartlog.FATAL, "main", "Invalid version format", map[string]any{"version": version})
	}

	args := argparse.NewParser("hello-actions", "A simple Go web application")

	host := args.String("i", "ip", &argparse.Options{
		Required: false,
		Help:     "Host address to bind to (default: 0.0.0.0)",
	})
	port := args.String("p", "port", &argparse.Options{
		Required: false,
		Help:     "Port to bind to (default: 8080)",
	})
	isDebug := args.Flag("d", "debug", &argparse.Options{
		Required: false,
		Help:     "Enable debug mode",
	})

	err := args.Parse(os.Args)
	if err != nil {
		logger.Log(dartlog.FATAL, "main", "Error parsing arguments", map[string]any{"error": err.Error()})
	}

	logger.Log(dartlog.INFO, "main", "Starting application...", nil)
	config, err := config.LoadConfig(*host, *port, *isDebug)
	if err != nil {
		logger.Log(dartlog.FATAL, "main", "Configuration error", map[string]any{"error": err.Error()})
	}
	logger.SetIsDebug(config.IsDebug)
	logger.Log(dartlog.INFO, "main", "Configuration loaded", map[string]any{"host": config.Host, "port": config.Port, "is_debug": config.IsDebug})

	app := app.New("hello-actions", utils.ExtractVersion(version), config, logger)

	if err := app.Run(); err != nil {
		logger.Log(dartlog.FATAL, "main", "Application error", map[string]any{"error": err.Error()})
	}
	logger.Log(dartlog.INFO, "main", "Application finished successfully", nil)

}
