package config

import (
	"testing"
)

func TestParseEnv(t *testing.T) {
	t.Setenv("SHELL", "zsh")
	t.Setenv("HOME", "/home/quentin")

	var cfg Config
	err := ParseEnv(&cfg)
	if err != nil {
		t.Fatalf("Failed to parse environment variables: %v", err)
	}

	if cfg.ShellType != "zsh" {
		t.Errorf("Expected ShellType 'zsh', got %s", cfg.ShellType)
	}
	if cfg.HistoryPath != "/home/quentin/.zsh_history" {
		t.Errorf("Expected HistoryPath '/home/quentin/.zsh_history', got %s", cfg.HistoryPath)
	}
}
