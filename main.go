package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"
)

type Config struct {
	Top      int
	Shell    string
	Histfile string
}

var usage = `Usage: %s [options]
Print top commands used from your shell history

Options:
`

func main() {
	var config Config
	config.Top = 5
	if err := parseEnv(&config); err != nil {
		log.Fatalf("Failed to parse environment variables: %v", err)
	}

	flag.Usage = func() {
		name := path.Base(os.Args[0])
		fmt.Fprintf(os.Stderr, usage, name)
		flag.PrintDefaults()
	}
	flag.Var(&Top{&config.Top}, "top", "how many top commands to return")
	flag.Parse()

	fmt.Println("Success!")
}
