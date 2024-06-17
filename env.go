package main

import (
	"fmt"
	"os"
	"strings"
)

func parseEnv(config *Config) error {
	const shellKey = "SHELL"
	const histfileKey = "HISTFILE"
	const xdgDataHomeKey = "XDG_DATA_HOME"

	shellPath := os.Getenv(shellKey)
	parts := strings.Split(shellPath, "/")
	shellName := parts[len(parts)-1]

	// Set history file according to current shell, if possible
	switch shellName {
	case "bash":
		config.Histfile = os.Getenv(histfileKey)
	case "zsh":
		config.Histfile = os.Getenv(histfileKey)
	case "fish":
		config.Histfile = os.Getenv(xdgDataHomeKey) + "/fish/fish_history"
	case "":
		fmt.Errorf("can't parse %s environment variable value", shellKey)
	default:
		fmt.Errorf("unrecognized shell %s", shellName)
	}

	config.Shell = shellName

	return nil
}
