package history

import (
	"bufio"
	"os"
	"strings"

	"github.com/quentinlintz/cmdtop/models"
)

type PowerShellParser struct{}

func (p *PowerShellParser) ParseHistory(filePath string) ([]models.Command, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	commandMap := make(map[string]int)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			// Split by whitespace and take the first part as the command
			parts := strings.Fields(line)
			if len(parts) > 0 {
				commandMap[parts[0]]++
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	var commands []models.Command
	for name, count := range commandMap {
		commands = append(commands, models.Command{Name: name, Count: count})
	}

	return commands, nil
} 