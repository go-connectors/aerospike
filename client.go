package aerospike

import (
	"errors"

	as "github.com/aerospike/aerospike-client-go"
	"github.com/aerospike/aerospike-client-go/types"
)

// Client encapsulates an Aerospike cluster.
// All database operations are available against this object.
type Client struct {
	*as.Client
	Config Config
}

// NewClient creates connection to Aerospike and returns Aerospike Client.
func NewClient(cfg Config) (*Client, error) {
	client, err := as.NewClient(cfg.Host, cfg.Port)
	if err != nil {
		return nil, err
	}

	return &Client{Client: client, Config: cfg}, nil
}

// IsErrorNotFound checks error for Aerospike ErrKeyNotFound.
func IsErrorNotFound(err error) bool {
	return errors.Is(err, types.ErrKeyNotFound)
}
