package runtimes

import (
	"sync"

	"github.com/Nerv-Core-Developers/nerv-provider-golang/shared"
)

type NervRuntime struct {
	currentRunning struct {
		Functions sync.Map
		Services  sync.Map
	}
	Node   *shared.NodeCfg
	Logger shared.LoggerInterface
}
