package history

import (
	"fmt"
	"log"
	"sort"

	"github.com/quentinlintz/cmdtop/config"
	"github.com/quentinlintz/cmdtop/models"
)

type Parser interface {
	ParseHistory(filePath string) ([]models.Command, error)
}

func PrintTopCommands(cfg config.Config, p Parser) {
	commands, err := p.ParseHistory(cfg.HistoryPath)
	if err != nil {
		log.Fatalf("Error parsing history: %v", err)
	}

	sort.Slice(commands, func(i, j int) bool {
		return commands[i].Count > commands[j].Count
	})

	fmt.Printf("Top %d commands:\n", cfg.Top)
	for i, cmd := range commands {
		if i >= cfg.Top {
			break
		}
		fmt.Printf("%d: %s (%d)\n", i+1, cmd.Name, cmd.Count)
	}
}
