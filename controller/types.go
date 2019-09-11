package controller

import "github.com/Nerv-Core-Developers/nerv-provider-golang/shared"

type Controller struct {
	Node   *shared.NodeCfg
	Logger shared.LoggerInterface
}

type Registry struct {
	Name string
	ID   string
}
