package workspace

import (
	"ai-terminal/internal/config"
	"ai-terminal/internal/logger"
	"ai-terminal/internal/types"
	. "ai-terminal/internal/types"
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/term"
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

var workspace []UserWorkspaceConfig

func createWorkspaceConfig() (bool, error) {

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

func loadWorkspaceConfig() (string, error) {
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

	err = json.Unmarshal(fileBytes, &workspace)

	if err != nil {
		logger.Error("Error loading config file: %v", err)
	}
	return "", nil
}

func ProcessWorkspaceConfig() error {

	file, err := createWorkspaceConfig()
	if err != nil {
		return err
	}
	if file {
		work, err := loadWorkspaceConfig()
		if err != nil {
			return err
		}

		if work == types.EMPTY_FILE {
			// create first workspaces
			createWorkSpace()
			promptWorkspace()
		}
	}
	return nil
}

func createWorkSpace() {
	logger.Info("Create new workspace")
	path, err := os.Getwd()

	if err != nil {
		logger.Error("Workspace Path error: %v", err)
	}

	logger.Info("Path: %s", path)
}

func promptWorkspace() (UserWorkspaceConfig, error) {

	reader := bufio.NewReader(os.Stdin)

	path, err := os.Getwd()

	if err != nil {
		logger.Error("Path is not recognised: %v", err)
	}

	fmt.Printf("\n==============================================\n")
	fmt.Println(ColorBlue, "           MY AI-Terminal            \n", ColorReset)
	fmt.Println(ColorMagneta, "=========== Welcome to AI-Terminal =============\n", ColorReset)
	fmt.Printf("==============================================\n")

	selectTrust()

	fmt.Println(BrightBlue, "\nProject Directory: ", path, ColorReset)

	fmt.Println("\nEnter the project name: ")

	projectName, err := reader.ReadString('\n')
	if err != nil {
		return UserWorkspaceConfig{}, err
	}

	projectName = strings.TrimSpace(projectName)

	logger.Info("project name: %s", projectName)

	return UserWorkspaceConfig{}, nil
}

func selectTrust() (TrustStatus, error) {
	options := []string{
		"No, I don't trust this folder",
		"Yes, I trust this folder",
	}

	selected := 0

	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		return Unknown, err
	}

	defer term.Restore(int(os.Stdin.Fd()), oldState)

	fmt.Print("\033[?25l")
	defer fmt.Print("\033[?25h")

	for {
		// Redraw menu
		fmt.Print("\033[5;1H")

		fmt.Println("Do you trust this folder?")
		fmt.Println()

		for i, option := range options {
			if selected == i {
				fmt.Printf("    \033[32m> %s\033[0m\n", option)
			} else {
				fmt.Printf("    %s\n", option)
			}
		}

		fmt.Println("\n  ↑ ↓ Navigate   Enter Select")

		var input [1]byte

		_, err := os.Stdin.Read(input[:])
		if err != nil {
			return Unknown, nil
		}

		switch input[0] {
		case 13, 10: // ENTER
			if selected == 0 {
				return UNTRUSTED, nil
			} else {
				return TRUSTED, nil
			}
		case 3:
			return Unknown, fmt.Errorf("Selection Cancelled")
		case 27:
			var seq [2]byte

			_, err := os.Stdin.Read(seq[:])
			if err != nil {
				return Unknown, nil
			}

			if seq[0] == '[' {
				switch seq[1] {
				case 'A':
					selected--
					if selected < 0 {
						selected = len(options) - 1
					}
				case 'B':
					selected++
					if selected >= len(options) {
						selected = 0
					}
				}
			}
		}

	}
}
