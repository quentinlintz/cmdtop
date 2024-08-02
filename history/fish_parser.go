package history

import (
	"bufio"
	"os"
	"strings"

	"github.com/quentinlintz/cmdtop/models"
)

type FishParser struct{}

func (p *FishParser) ParseHistory(filePath string) ([]models.Command, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	commandMap := make(map[string]int)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		after, found := strings.CutPrefix(line, "- cmd: ")
		if found {
			command := strings.Fields(after)
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
