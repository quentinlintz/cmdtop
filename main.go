package main

import (
	"fmt"
	"log"

	"github.com/quentinlintz/cmdtop/config"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	fmt.Printf("Config: %+v\n", cfg)
}
