package history

import (
	"path/filepath"
	"testing"

	"github.com/quentinlintz/cmdtop/models"
)

var expectedCommands = []models.Command{
	{Name: "ls", Count: 2},
	{Name: "cd", Count: 1},
	{Name: "git", Count: 2},
}

func TestZshParseHistoryContents(t *testing.T) {
	testFile := filepath.Join("testdata", "zsh_history")
	parser := &ZshParser{}
	testParser(t, parser, testFile, expectedCommands)
}

func TestZshParseHistoryFileNotFound(t *testing.T) {
	testFile := filepath.Join("testdata", "not_found")
	parser := &ZshParser{}
	_, err := parser.ParseHistory(testFile)
	if err == nil {
		t.Fatal("Expected error when parsing non-existent file, but got none")
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

	for i, cmd := range commands {
		if cmd.Name != expected[i].Name || cmd.Count != expected[i].Count {
			t.Errorf("Expected command %v, got %v", expected[i], cmd)
		}
	}
}
