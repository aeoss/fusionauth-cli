package batcher

import (
	"github.com/aeoss/fusionauth-cli/pkg/client"
)

var registeredBatchers = make([]Batcher, 0)

// Batcher implements the necessary items to manage a type of configuration within
// the FusionAuth instance.
type Batcher interface {
	ResolveData
	Exporter
	Importer

	// GetName returns the name of the Batcher.
	GetName() string
}

// ResolveData is responsible for resolving the files and/or directories used
// by the
type ResolveData interface {

	// GetKinds returns the Kinds that are supported by this Batcher.
	GetKinds() []string

	// ConvertSpec converts the provided bytes into a statically typed struct
	// specific to this Batcher.
	ConvertSpec(data []byte) (interface{}, error)

	// Save saves the array of specs to a local file in an oppinionated directory.
	Save(spec []interface{}) error
}

// Exporter fetches data from FusionAuth.
type Exporter interface {
	ExportAll(client *client.Client) ([]interface{}, error)
	ExportByID(client *client.Client, id string) (interface{}, error)
	ExportByName(client *client.Client, name string) (interface{}, error)
}

// Importer is an interface designed for importing data into FusionAuth
// and all methods should always overwrite any existing data that matches.
type Importer interface {
	ImportAll(client *client.Client, spec []interface{}) error
	ImportByID(client *client.Client, id string, spec interface{}) error
	ImportByName(client *client.Client, name string, spec interface{}) error
}

// Client is the Batcher client used to manipulate the data on the
// FusionAuth side.
type Client struct {
	Dir        string
	Batchers   map[string]Batcher
	HTTPClient *client.Client
}

// GetBatcher returns a batcher by name.
func (c *Client) GetBatcher(name string) Batcher {
	if c.Batchers[name] != nil {
		return c.Batchers[name]
	}

	return nil
}
