package shared

type YamlPackageMetaData struct {
	Version  int    `yaml:"version"`
	Owner    string `yaml:"owner"`
	OwnerSig string `yaml:"owner-sig"`
	Files    struct {
		Functionals []string `yaml:"functionals"`
		Services    []string `yaml:"services"`
		Assets      []string `yaml:"assets"`
	} `yaml:"files"`
}

type YamlConfigure struct {
	Version       int            `yaml:"version"`
	NetworkID     int            `yaml:"network-id"`
	Modules       []string       `yaml:"modules"`
	Identity      string         `yaml:"identity"`
	Owner         string         `yaml:"owner"`
	DataDir       string         `yaml:"datadir"`
	Debug         bool           `yaml:"debug"`
	DebugRuntime  bool           `yaml:"debug-runtime"`
	Network       YamlNetworkCfg `yaml:"network"`
	Policy        YamlPolicyCfg  `yaml:"policy"`
	Logger        YamlLogCfg     `yaml:"log-config"`
	RegistriesDir string         `yaml:"registries-dir"`
}

type YamlNetworkCfg struct {
	Port           int      `yaml:"port"`
	APIPort        int      `yaml:"api-port"`
	BroadcastAddrs []string `yaml:"broadcast-addrs"`
	InitProviders  []string `yaml:"init-peers"`
	LinkNodes      []string `yaml:"link-nodes"`
}
type YamlPolicyCfg struct {
	AssetPolicy   YamlAssetPolicy   `yaml:"asset"`
	RuntimePolicy YamlRuntimePolicy `yaml:"runtime"`
	NetworkPolicy YamlNetworkPolicy `yaml:"network"`
}

type YamlAssetPolicy struct {
	MaxAssetSize uint64 `yaml:"max-asset-size"`
	Retriever    struct {
		MaxConnPerFile            int `yaml:"max-conn-per-file"`
		MaxConconcurrentRetriever int `yaml:"max-concurrent-retriever"`
	} `yaml:"retriever"`
	Uploader struct {
		MaxConcurrentUploader int `yaml:"max-concurrent-uploader"`
	} `yaml:"uploader"`
}

type YamlRuntimePolicy struct {
	MaxRunningService int `yaml:"max-running-service"`
	MaxConCurrentCall struct {
		Service  int `yaml:"service"`
		Function int `yaml:"function"`
	} `yaml:"max-concurrent-calls"`
	AllowRuntimeTypes []string `yaml:"allow-runtime-types"`
}

type YamlNetworkPolicy struct {
	Providers struct {
		Min     int `yaml:"min"`
		Max     int `yaml:"max"`
		Latency int `yaml:"latency"`
	} `yaml:"providers"`
	Clients struct {
		Max     int `yaml:"max"`
		Latency int `yaml:"latency"`
	} `yaml:"clients"`
}

type YamlLogCfg struct {
	Logger struct {
		Endpoint     string `yaml:"endpoint"`
		EndpointType string `yaml:"endpoint-type"`
	} `yaml:"logger"`
	Functional struct {
		Enable   bool   `yaml:"enable"`
		Endpoint string `yaml:"endpoint"`
	} `yaml:"functional"`
	Service struct {
		Enable   bool   `yaml:"enable"`
		Endpoint string `yaml:"endpoint"`
	} `yaml:"service"`
}
