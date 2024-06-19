package config

import (
	"fmt"
	"os"
	"strings"
)

func ParseEnv(config *Config) error {
	const shellKey = "SHELL"

	shellPath := os.Getenv(shellKey)
	parts := strings.Split(shellPath, "/")
	shellName := parts[len(parts)-1]

	home, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("can't get home directory")
	}

	// Set history file according to current shell, if possible
	switch shellName {
	case "zsh":
		config.HistoryPath = home + "/.zsh_history"
	case "":
		return fmt.Errorf("can't parse %s environment variable value", shellKey)
	default:
		return fmt.Errorf("unsupported shell %s", shellName)
	}

	config.ShellType = shellName

	return nil
}
