package main

import (
	"ai-terminal/internal/config"
	"ai-terminal/internal/logger"
	"ai-terminal/internal/workspace"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		logger.Fatal("Env Load error %s", err)
	}

	config.Load()

	if err != nil {
		logger.Error("%v", err)
	}

	workspace.ProcessWorkspaceConfig()

	// logger.Debug("AI-Terminal Start")

	// logger.Info("Info About AI Terminal")
}
