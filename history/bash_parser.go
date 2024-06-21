package history

import (
	"bufio"
	"os"
	"strings"

	"github.com/quentinlintz/cmdtop/models"
)

type BashParser struct{}

func (p *BashParser) ParseHistory(filepath string) ([]models.Command, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	commandMap := make(map[string]int)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")
		command := parts[0]
		if len(command) > 0 {
			commandMap[command]++
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
