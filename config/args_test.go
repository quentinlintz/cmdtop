package config

import (
	"testing"
)

func TestTopSet(t *testing.T) {
	tests := []struct {
		input    string
		expected int
		wantErr  bool
	}{
		{"10", 10, false},
		{"1", 1, false},
		{"100", 100, false},
		{"0", 0, true},
		{"101", 0, true},
		{"abc", 0, true},
		{"-1", 0, true},
	}

	for _, tt := range tests {
		var topVal int
		top := Top{&topVal}
		err := top.Set(tt.input)

		if tt.wantErr {
			if err == nil {
				t.Errorf("expected error for input %q, but got none", tt.input)
			}
		} else {
			if err != nil {
				t.Errorf("did not expect error for input %q, but got %v", tt.input, err)
			}
			if *top.val != tt.expected {
				t.Errorf("expected %d for input %q, but got %d", tt.expected, tt.input, *top.val)
			}
		}
	}
}
