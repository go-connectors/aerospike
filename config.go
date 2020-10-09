package client

import (
	"errors"
	"fmt"
)

// Config validation errors.
var (
	// ErrConfigValidation is general config validation error message.
	ErrConfigValidation = errors.New("aerospike config validation error")

	ErrEmptyHost      = fmt.Errorf("%w: host is empty", ErrConfigValidation)
	ErrEmptyPort      = fmt.Errorf("%w: port is empty", ErrConfigValidation)
	ErrEmptyNamespace = fmt.Errorf("%w: namespace is empty", ErrConfigValidation)
)

// Config contains configuration Aerospike data.
type Config struct {
	Host      string `json:"host" yaml:"host"`
	Port      int    `json:"port" yaml:"port"`
	Namespace string `json:"namespace" yaml:"namespace"`
}

// Validate checks required fields.
func (c *Config) Validate() error {
	if c.Host == "" {
		return ErrEmptyHost
	}

	if c.Port == 0 {
		return ErrEmptyPort
	}

	if c.Namespace == "" {
		return ErrEmptyNamespace
	}

	return nil
}
