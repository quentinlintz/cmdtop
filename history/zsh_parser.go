package history

import (
	"bufio"
	"os"
	"strings"

	"github.com/quentinlintz/cmdtop/models"
)

type ZshParser struct{}

func (p *ZshParser) ParseHistory(filePath string) ([]models.Command, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	commandMap := make(map[string]int)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		var commandText string
		if strings.HasPrefix(line, ":") {
			parts := strings.SplitN(line, ";", 2)
			if len(parts) == 2 {
				commandText = parts[1]
			}
		} else {
			commandText = line
		}
		if commandText != "" {
			command := strings.Fields(commandText)
			if len(command) > 0 {
				commandMap[command[0]]++
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
