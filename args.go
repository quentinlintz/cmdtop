package main

import (
	"fmt"
	"strconv"
)

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
