package main

import (
	"github.com/dart-chain/hello-actions/internal/app"
	"github.com/dart-chain/hello-actions/internal/utils"
	"gitlab.com/tim_de/dartlog/v2/dartlog"
)

var (
	version string
)

func main() {
	logger := dartlog.New("hello-actions", true, true)

	logger.Log(dartlog.INFO, "main", "Starting application...", nil)

	if !utils.IsVersionValid(version) {
		logger.Log(dartlog.FATAL, "main", "Invalid version format", map[string]any{"version": version})
	}

	app := app.New("hello-actions", utils.ExtractVersion(version), logger)

	if err := app.Run(); err != nil {
		logger.Log(dartlog.FATAL, "main", "Application error", map[string]any{"error": err.Error()})
	}
	logger.Log(dartlog.INFO, "main", "Application finished successfully", nil)

}
