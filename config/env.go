package config

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func ParseEnv(config *Config) error {
	const shellKey = "SHELL"

	shellPath := os.Getenv(shellKey)
	var shellName string

	// Special handling for Windows PowerShell
	if runtime.GOOS == "windows" && strings.Contains(strings.ToLower(os.Getenv("PSModulePath")), "windowspowershell") {
		shellName = "powershell"
	} else if shellPath != "" {
		// Unix-like shell detection
		shellName = filepath.Base(shellPath)
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("can't get home directory")
	}

	// Set history file according to current shell, if possible
	switch shellName {
	case "zsh":
		config.HistoryPath = filepath.Join(home, ".zsh_history")
	case "bash":
		config.HistoryPath = filepath.Join(home, ".bash_history")
	case "fish":
		config.HistoryPath = filepath.Join(home, ".local", "share", "fish", "fish_history")
	case "powershell":
		config.HistoryPath = filepath.Join(home, "AppData", "Roaming", "Microsoft", "Windows", "PowerShell", "PSReadLine", "ConsoleHost_history.txt")
	case "":
		return fmt.Errorf("can't parse %s environment variable value", shellKey)
	default:
		return fmt.Errorf("unsupported shell %s", shellName)
	}

	config.ShellType = shellName

	return nil
}
