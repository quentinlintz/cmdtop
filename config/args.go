package config

import (
	"flag"
	"fmt"
	"os"
	"path"
	"strconv"
)

var usage = `Usage: %s [options]
Print top commands used from your shell history

Options:
`

type Config struct {
	ShellType   string
	HistoryPath string
	TopN        int
}

type Top struct {
	val *int
}

func (t *Top) String() string {
	if t.val == nil {
		return ""
	}

	return fmt.Sprintf("%d", *t.val)
}

func (t *Top) Set(val string) error {
	n, err := strconv.Atoi(val)
	if err != nil {
		return fmt.Errorf("bad number: %v", err)
	}

	if n <= 0 || n > 100 {
		return fmt.Errorf("top %d out of range [1:100]", n)
	}

	*t.val = n

	return nil
}

func Load() (Config, error) {
	// Create config and set default top commands to 5
	var config Config
	config.TopN = 5

	flag.Usage = func() {
		name := path.Base(os.Args[0])
		fmt.Fprintf(os.Stderr, usage, name)
		flag.PrintDefaults()
	}
	flag.Var(&Top{&config.TopN}, "top", "how many top commands to return")
	flag.Parse()

	if err := parseEnv(&config); err != nil {
		return config, fmt.Errorf("failed to parse environment variables")
	}

	return config, nil
}
