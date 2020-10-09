package client

import (
	"testing"
)

func TestConfig_Validate(t *testing.T) {
	var cases = []struct {
		name   string
		config Config
		err    error
	}{
		{name: "valid", config: Config{Host: "127.0.0.1", Port: 3000, Namespace: "test"}, err: nil},
		{name: "empty host", config: Config{Port: 3000, Namespace: "test"}, err: ErrEmptyHost},
		{name: "empty port", config: Config{Host: "127.0.0.1"}, err: ErrEmptyPort},
		{name: "empty namespace", config: Config{Host: "127.0.0.1", Port: 3000}, err: ErrEmptyNamespace},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.config.Validate()
			if err != tt.err {
				t.Errorf("got %q, want %q", err, tt.err)
			}
		})
	}
}
