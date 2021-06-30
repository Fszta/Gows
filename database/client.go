package database

import (
	"com.github/Fszta/gows/pkg/dag"
)

var DbClient Client

// Database Client interface
type Client interface {
	init() error
	SaveDagConfig(*dag.DagConfig) error
	GetAllDags() []dag.DagConfig
	GetDag(dagUUID string) (*dag.DagConfig, error)
	RemoveDag(dagUUID string) error
	Close()
}

// Create new db client
func NewClient() (Client, error) {
	client := &BoltClient{}
	err := client.init()

	if err != nil {
		return nil, err
	}
	return client, nil
}
