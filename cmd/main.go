package main

import "ai-terminal/internal/logger"

func main() {
	logger.SetLogLevel(logger.LevelDebug)
	logger.Debug("AI-Terminal Start")

	logger.Info("Info About AI Terminal")
}
