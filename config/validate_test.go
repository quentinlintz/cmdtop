package config

import (
	"testing"
)

func TestValidateConfig(t *testing.T) {
	tests := []struct {
		name    string
		config  Config
		wantErr bool
	}{
		{"ValidTop10", Config{Top: 10}, false},
		{"ValidTop1", Config{Top: 1}, false},
		{"ValidTop100", Config{Top: 100}, false},
		{"InvalidTop0", Config{Top: 0}, true},
		{"InvalidTop101", Config{Top: 101}, true},
		{"NegativeTop", Config{Top: -5}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateConfig(&tt.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateConfig() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
