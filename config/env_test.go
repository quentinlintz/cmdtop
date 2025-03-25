package config

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestParseEnv(t *testing.T) {
	// Save original environment
	originalPSModulePath := os.Getenv("PSModulePath")
	originalShell := os.Getenv("SHELL")
	defer func() {
		os.Setenv("PSModulePath", originalPSModulePath)
		os.Setenv("SHELL", originalShell)
	}()

	tests := []struct {
		name            string
		setupEnv        func()
		expectedShell   string
		expectedHistory string
	}{
		{
			name: "Unix ZSH",
			setupEnv: func() {
				os.Setenv("SHELL", "/bin/zsh")
				os.Setenv("PSModulePath", "")
			},
			expectedShell:   "zsh",
			expectedHistory: ".zsh_history",
		},
		{
			name: "Windows PowerShell",
			setupEnv: func() {
				os.Setenv("PSModulePath", "WindowsPowerShell")
				os.Setenv("SHELL", "")
			},
			expectedShell:   "powershell",
			expectedHistory: filepath.Join("AppData", "Roaming", "Microsoft", "Windows", "PowerShell", "PSReadLine", "ConsoleHost_history.txt"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupEnv()

			var cfg Config
			err := ParseEnv(&cfg)
			if err != nil {
				t.Fatalf("Failed to parse environment variables: %v", err)
			}

			if cfg.ShellType != tt.expectedShell {
				t.Errorf("Expected ShellType '%s', got %s", tt.expectedShell, cfg.ShellType)
			}

			if !strings.HasSuffix(cfg.HistoryPath, tt.expectedHistory) {
				t.Errorf("Expected HistoryPath to end with '%s', got %s", tt.expectedHistory, cfg.HistoryPath)
			}
		})
	}
}
