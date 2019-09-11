package shared

import (
	"archive/zip"

	"github.com/Nerv-Core-Developers/nerv-provider-golang/shared/nna"
	"github.com/Nerv-Core-Developers/nerv-provider-golang/shared/noid"
	"github.com/Nerv-Core-Developers/nerv-provider-golang/shared/schema"
)

type NodeCfg struct {
	Info         *NodeInfo
	Network      NetworkInterface
	Controler    ControllerInterface
	Runtime      RuntimeInterface
	Pubsub       PubsubInterface
	DBCache      DatabaseCacheInterface
	DBPersistent DatabasePersistentInterface
}

type NodeInfo struct {
	ID  string
	Cfg *YamlConfigure
}

type PackageData struct {
	MetaData YamlPackageMetaData
	Files    map[string]*zip.File
}

type FunctionInfo struct {
	ID          noid.Identifier
	RuntimeType schema.RuntimeType
	Local       bool
	Owner       string
	OwnerSig    string
}

type NetMsg struct {
	ReceiveFrom string
	SenderType  nna.NodeType
	MsgData     schema.Msg
}

type AssetLink struct {
	// Addresses  []string
	// ProviderID string
	AssetID noid.Identifier
	Size    int64
}

type ProviderInfo struct {
	ProviderID string
	Addresses  []string
	Version    NodeVersion
	Owner      string
	Reachable  bool
}

type RequestInfo struct {
	RequestorType schema.NodeType
	Type          schema.MsgType
	Requestor     string
	RequestCtx    interface{}
}

type RespondMsg struct {
	Msg    schema.Msg
	Sender string
}

type NodeVersion struct {
	Controller int
	Network    int
	Runtime    int
}

type DBPatch struct {
	Key  string
	Data interface{}
}

type RuntimeOutput struct {
	Status bool
	Result string
	Logs   []string
}
