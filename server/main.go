package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/gdamore/tcell"
	"github.com/Nerv-Core-Developers/nerv-provider-golang/controller"
	"github.com/Nerv-Core-Developers/nerv-provider-golang/database"
	"github.com/Nerv-Core-Developers/nerv-provider-golang/logger"
	"github.com/Nerv-Core-Developers/nerv-provider-golang/network"
	"github.com/Nerv-Core-Developers/nerv-provider-golang/pubsub"
	"github.com/Nerv-Core-Developers/nerv-provider-golang/runtimes"
	"github.com/Nerv-Core-Developers/nerv-provider-golang/shared"
	uuid "github.com/satori/go.uuid"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	configFilePtr := flag.String("config", "./config.yaml", "path to config file")
	flag.Parse()

	fmt.Println("[nerv-provider]", "["+fmt.Sprintf("CTRL: %d | NNET: %d | NRT: %d", controller.Version, network.Version, runtimes.Version)+"]", "\n")

	fmt.Println("Loading config file...")
	loadConfig, err := loadConfig(*configFilePtr)
	if err != nil {
		panic(err)
	}
	if err := logger.InitLogger(loadConfig.DataDir, &loadConfig.Logger); err != nil {
		panic(err.Error())
	}
	if err = systemCheck(loadConfig); err != nil {
		panic("System doesn't satify requirement to run your config " + err.Error())
	}
	initStorageFolder(loadConfig.DataDir)
	fmt.Println("Initializing TUI...")

	err = SetupTUI(exitEventHandler)
	if err != nil {
		panic(err)
	}
	time.AfterFunc(1*time.Second, func() {
		go logger.StartLogRecorder(MainLog)
		go startNode(loadConfig)
	})
	StartTUI()
}

func startNode(cfg *shared.YamlConfigure) {
	var nodeID string
	if cfg.Identity != "" {
		nodeID = uuid.NewV3(uuid.NamespaceX500, cfg.Identity).String()
	} else {
		nodeID = uuid.NewV4().String()
	}
	nodeCfg := shared.NodeCfg{
		Info: &shared.NodeInfo{
			ID:  nodeID,
			Cfg: cfg,
		},
		DBCache:      &database.DBcache{Logger: &logger.Logger{}},
		DBPersistent: &database.DBpersistent{Logger: &logger.Logger{}},
		Pubsub:       &pubsub.PubsubService{},
	}
	if err := nodeCfg.Pubsub.Init(); err != nil {
		panic(err.Error())
	}
	if err := nodeCfg.DBCache.Init("cachedb", nodeCfg.Info.Cfg.DataDir); err != nil {
		panic(err.Error())
	}
	if err := nodeCfg.DBPersistent.Init("persistentdb", nodeCfg.Info.Cfg.DataDir); err != nil {
		panic(err.Error())
	}
	networkInstance := network.Network{
		Node:   &nodeCfg,
		Logger: &logger.Logger{},
	}
	controllerInstance := controller.Controller{
		Node:   &nodeCfg,
		Logger: &logger.Logger{},
	}
	runtimeInstance := runtimes.NervRuntime{
		Node:   &nodeCfg,
		Logger: &logger.Logger{},
	}

	nodeCfg.Network = &networkInstance
	nodeCfg.Controler = &controllerInstance
	nodeCfg.Runtime = &runtimeInstance

	if err := nodeCfg.Network.Init(); err != nil {
		panic(err.Error())
	}
	if err := nodeCfg.Controler.Init(); err != nil {
		panic(err.Error())
	}
	if err := nodeCfg.Runtime.Init(); err != nil {
		panic(err.Error())
	}
}

func initStorageFolder(dataDir string) {
	fmt.Println("Initializing storage...")
	if err := os.MkdirAll(filepath.Dir(dataDir), os.ModePerm); err != nil {
		panic(err)
	}
	if err := os.MkdirAll(filepath.Dir(dataDir+"/storage"), os.ModePerm); err != nil {
		panic(err)
	}
	return
}

func exitEventHandler(event *tcell.EventKey) *tcell.EventKey {
	if event.Key() == tcell.KeyCtrlC {
		App.Stop()

		fmt.Println("\nStopping nerv-node...")
		os.Exit(1)
	}
	return event
}

func systemCheck(cfg *shared.YamlConfigure) error {
	return nil
}
