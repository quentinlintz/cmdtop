package history

import (
	"path/filepath"
	"sort"
	"testing"

	"github.com/quentinlintz/cmdtop/models"
)

var expectedCommands = []models.Command{
	{Name: "ls", Count: 2},
	{Name: "cd", Count: 1},
	{Name: "git", Count: 2},
}

func TestParseHistory(t *testing.T) {
	tests := []struct {
		name        string
		parser      Parser
		historyFile string
	}{
		{"Zsh", &ZshParser{}, "zsh_history"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testFile := filepath.Join("testdata", tt.historyFile)
			testParser(t, tt.parser, testFile, expectedCommands)
		})
	}
}

func TestHistoryFileNotFound(t *testing.T) {
	tests := []struct {
		name   string
		parser Parser
	}{
		{"Zsh", &ZshParser{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testFile := filepath.Join("testdata", "not_found")
			_, err := tt.parser.ParseHistory(testFile)
			if err == nil {
				t.Fatal("Expected error when parsing non-existent file, but got none")
			}
		})
	}
}

func testParser(t *testing.T, parser Parser, testFile string, expected []models.Command) {
	commands, err := parser.ParseHistory(testFile)
	if err != nil {
		t.Fatalf("ParseHistory failed: %v", err)
	}

	if len(commands) != len(expected) {
		t.Fatalf("Expected %d commands, got %d", len(expected), len(commands))
	}

	sort.Slice(commands, func(i, j int) bool {
		return commands[i].Name < commands[j].Name
	})

	sort.Slice(expected, func(i, j int) bool {
		return expected[i].Name < expected[j].Name
	})

	for i, cmd := range commands {
		if cmd.Name != expected[i].Name || cmd.Count != expected[i].Count {
			t.Errorf("Expected command %v, got %v", expected[i], cmd)
		}
	}
}
