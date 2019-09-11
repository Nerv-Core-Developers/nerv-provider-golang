package network

import (
	"github.com/Nerv-Core-Developers/nerv-provider-golang/shared"
	"github.com/Nerv-Core-Developers/nerv-provider-golang/shared/nna"
	"github.com/Nerv-Core-Developers/nerv-provider-golang/shared/schema"
)

func (nw *Network) Init() error {
	nw.Logger.Init("network")
	nw.Logger.Log("Init network...")
	return nil
}

func (nw *Network) Start() error {
	return nil
}
func (nw *Network) BroadcastMsgToProviders(msg *shared.NetMsg) error {
	return nil
}
func (nw *Network) GetProviderAddresses(string) []string {
	return nil
}
func (nw *Network) GetProviderList() []shared.ProviderInfo {
	var result []shared.ProviderInfo
	return result
}
func (nw *Network) SendMsgToClient(*schema.Msg, string) error {
	return nil
}
func (nw *Network) SendMsgToProvider(*schema.Msg, string) error {
	return nil
}
func (nw *Network) GetAvailableNodeTypes() []nna.NodeType {
	return nil
}
func (nw *Network) GetAvailableConnectionTypes() []nna.ConnectionType {
	return nil
}
func (nw *Network) GetAvailableConnectionProtocolTypes() []nna.ConnectionProtocolType {
	return nil
}
func (nw *Network) Version() int {
	return Version
}
