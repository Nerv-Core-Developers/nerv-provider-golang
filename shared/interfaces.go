package shared

import (
	"mime/multipart"
	"net/http"

	"github.com/Nerv-Core-Developers/nerv-provider-golang/shared/nna"
	"github.com/Nerv-Core-Developers/nerv-provider-golang/shared/noid"
	"github.com/Nerv-Core-Developers/nerv-provider-golang/shared/schema"
)

type DatabaseCacheInterface interface {
	Init(dbname string, dbpath string) error
	Export(path string) error
	Put(key string, value interface{}, timeout int, overwrite bool) error
	Batch([]DBPatch) error
	Get(key string) (interface{}, int, error)
	Delete(key string) error
}

type DatabasePersistentInterface interface {
	Init(dbname string, dbpath string) error
	Export(path string) error
	Put(key string, value []byte, overwrite bool) error
	Batch([]DBPatch) error
	Get(key string) ([]byte, error)
	Delete(key string) error
}

type ControllerInterface interface {
	Init() error
	GetFunctionInfo(ID noid.Identifier) (FunctionInfo, error)
	GetAssetLink(ID noid.Identifier) (string, error)
	DeployPkg(pkg *multipart.FileHeader) error
	APIHandler(*http.Request, []string) ([]byte, error)
	GetNodeVersion() NodeVersion
	HandleMessage(msg *NetMsg) error
	Version() int
}

type NetworkInterface interface {
	Init() error
	BroadcastMsgToProviders(*NetMsg) error
	GetProviderAddresses(string) []string
	GetProviderList() []ProviderInfo
	SendMsgToClient(*schema.Msg, string) error
	SendMsgToProvider(*schema.Msg, string) error
	GetAvailableNodeTypes() []nna.NodeType
	GetAvailableConnectionTypes() []nna.ConnectionType
	GetAvailableConnectionProtocolTypes() []nna.ConnectionProtocolType
	Version() int
}

type RuntimeInterface interface {
	Init() error
	RunFunction(fnLink string, fnType schema.RuntimeType, data []byte) (*RuntimeOutput, error)
	StartService() error
	Version() int
}

type PubsubInterface interface {
	Init() error
	ProcessCmd(buf []byte, msgID string, subscriber string, subscriberCh chan interface{})
}

type LoggerInterface interface {
	Init(packagename string) error
	Log(string)
	Debug(string)
	Fatal(error)
}
