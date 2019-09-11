package service

import "github.com/Nerv-Core-Developers/nerv-provider-golang/shared"

type ServiceInstance struct {
	Type int
	Name string
	ID   string
	API  struct {
		DBcache      shared.DatabaseInterface
		DBpersistent shared.DatabaseInterface
		Controller   shared.ControllerInterface
		Network      shared.NetworkInterface
	}
	OutputCh chan []byte
}

func (sv *ServiceInstance) Init() {

}

func (sv *ServiceInstance) Run() {

}
