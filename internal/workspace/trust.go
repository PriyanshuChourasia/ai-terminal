package workspace

import (
	"ai-terminal/internal/config"
	"ai-terminal/internal/logger"
	"ai-terminal/internal/types"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type TrustStatus string

const (
	TRUSTED   TrustStatus = "trusted"
	UNTRUSTED TrustStatus = "untrusted"
	Unknown   TrustStatus = "unknown"
)

type Workspace struct {
	Path     string
	Username string
	Trust    TrustStatus
}

func CheckForWorkspaceConfig() (bool, error) {

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return false, fmt.Errorf("Error for Home DIR %v", err)
	}

	configData := homeDir + "/" + config.Env.ConfigDirName + "/" + config.Env.ConfigFileName

	fileInfo, err := os.Stat(configData)

	if err != nil {
		return false, fmt.Errorf("File Exist error %v", err)
	}

	if !fileInfo.IsDir() {
		return true, nil
	}

	folderDir := config.Env.ConfigDirName
	configFile := config.Env.ConfigFileName
	if folderDir == "" {
		return false, fmt.Errorf("No folder name in config")
	}
	dir := filepath.Join(homeDir, folderDir)

	info, err := os.Stat(dir)

	if err != nil {
		if os.IsNotExist(err) {
			if err := os.MkdirAll(dir, 0755); err != nil {
				logger.Error("Error creating folder at path")
				return false, fmt.Errorf("Error creating folder at path %v", err)
			}
		} else {
			return false, fmt.Errorf("Error checking folder %v", err)
		}
	}

	if configFile == "" {
		return false, fmt.Errorf("No config file is configured")
	}

	configFilePath := filepath.Join(dir, configFile)

	file, err := os.Create(configFilePath)
	if err != nil {
		return false, err
	}

	defer file.Close()

	info, err = os.Stat(configFilePath)

	if !info.IsDir() {
		return true, nil
	} else {
		return false, fmt.Errorf("File Path exists but it is not a file or folder")
	}
}

func LoadWorkspaceConfig() (string, error) {
	homeDir, err := os.UserHomeDir()

	if err != nil {
		return "", fmt.Errorf("Error checking home directory %v", err)
	}

	configFilePath := homeDir + "/" + config.Env.ConfigDirName + "/" + config.Env.ConfigFileName

	fileBytes, err := os.ReadFile(configFilePath)

	if err != nil {
		logger.Error("File reading error %v", err)
	}

	if len(fileBytes) == 0 {
		return types.EMPTY_FILE, nil
	}

	var workspace UserWorkspaceConfig

	err = json.Unmarshal(fileBytes, &workspace)

	if err != nil {
		logger.Error("Error loading config file: %v", err)
	}

	logger.Info("User %s", workspace.Username)

	logger.Info("HOME %s", configFilePath)
	return "", nil
}

func ProcessWorkspaceConfig() {

}
