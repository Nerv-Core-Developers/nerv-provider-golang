package network

import (
	"sync"

	"github.com/Nerv-Core-Developers/nerv-provider-golang/shared"
	"github.com/Nerv-Core-Developers/nerv-provider-golang/shared/nna"
	"github.com/Nerv-Core-Developers/nerv-provider-golang/shared/schema"
)

type Network struct {
	connectionList connectionList
	Node           *shared.NodeCfg
	Logger         shared.LoggerInterface
}

type Stream struct {
	stream  interface{}
	Address nna.NNA
	sync.Mutex
}

type providerConnection struct {
	streams  []*Stream
	MsgOutCh chan []byte
	MsgInCh  chan []byte
	netCtrl  *Network
	Info     *shared.ProviderInfo
	sync.RWMutex
}

type clientConnection struct {
	streams   []*Stream
	ID        string
	Namespace string
	MsgOutCh  chan []byte
	MsgInCh   chan []byte
	NetCtrl   *Network
}

type connectionList struct {
	Providers    map[string]*providerConnection
	providerList sync.RWMutex
	Clients      map[string]*clientConnection
	clientList   sync.RWMutex
}

func (list *connectionList) CheckExist(ID string, ntype schema.NodeType) bool {
	if ntype == schema.NodeTypeProvider {
		list.providerList.RLock()
		if _, ok := list.Providers[ID]; ok {
			list.providerList.RUnlock()
			return true
		}
		list.providerList.RUnlock()
	} else {
		list.clientList.RLock()
		if _, ok := list.Clients[ID]; ok {
			list.clientList.RUnlock()
			return true
		}
		list.clientList.RUnlock()
	}

	return false
}
