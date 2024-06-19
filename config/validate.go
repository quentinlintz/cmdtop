package config

import (
	"fmt"
)

type Config struct {
	ShellType   string
	HistoryPath string
	Top         int
	ShowVersion bool
	ShowLicense bool
}

func ValidateConfig(config *Config) error {
	// Top should be between 1 and 100
	if config.Top <= 0 || config.Top > 100 {
		return fmt.Errorf("top %d out of range [1:100]", config.Top)
	}

	return nil
}
