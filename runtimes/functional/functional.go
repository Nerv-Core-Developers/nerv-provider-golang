package functional

import (
	"github.com/naokichau/nerv-provider-golang/shared"
	"github.com/naokichau/nerv-provider-golang/shared/schema"
)

type FunctionInstance struct {
	Type int
	Name string
	ID   string
	API  struct {
		DBcache      shared.DatabaseCacheInterface
		DBpersistent shared.DatabasePersistentInterface
		Controller   shared.ControllerInterface
		Network      shared.NetworkInterface
	}
}

func (fn *FunctionInstance) Init() error {
	return nil
}

func (fn *FunctionInstance) Exec(data string) (*schema.RuntimeOutput, error) {
	return nil, nil
}
