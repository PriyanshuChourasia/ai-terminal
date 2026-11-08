package main

import (
	"ai-terminal/internal/config"
	"ai-terminal/internal/logger"
	"ai-terminal/internal/types"
	"ai-terminal/internal/workspace"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		logger.Fatal("Env Load error %s", err)
	}

	config.Load()

	file, err := workspace.CheckForWorkspaceConfig()
	if err != nil {
		logger.Error("%v", err)
	}

	if file {
		logger.Info("Path creation done")
		data, err := workspace.LoadWorkspaceConfig()
		if err != nil {
			logger.Error("Error loading workspace: %v", err)
		}
		if data == types.EMPTY_FILE {
			logger.Info("File is empty")
		}
	}

	// workspace.CheckForWorkspace()

	// logger.Debug("AI-Terminal Start")

	// logger.Info("Info About AI Terminal")
}
