package logger

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/Nerv-Core-Developers/nerv-provider-golang/shared"
)

type logConfig struct {
	DataDir      string
	EndpointType string
	Endpoint     string
	LogDebug     bool
}

var logCfg logConfig
var logStream chan string

func InitLogger(dataDir string, loggerCfg *shared.YamlLogCfg) error {
	logCfg = logConfig{
		DataDir:      dataDir,
		EndpointType: loggerCfg.Logger.EndpointType,
		Endpoint:     loggerCfg.Logger.Endpoint,
	}
	return nil
}

func StartLogRecorder(v io.Writer) error {
	flog, err := os.OpenFile(logCfg.DataDir+"/"+logCfg.Endpoint, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		println("error opening file: %v", err)
		return err
	}
	logStream = make(chan string)
	defer func() {
		close(logStream)
		flog.Close()
	}()
	for {
		select {
		case log := <-logStream:
			w := bufio.NewWriter(flog)
			fmt.Fprintln(w, log)
			w.Flush()
			highlight := string([]rune(log)[:3])
			switch highlight {
			case "LOG":
				highlight = "[yellow:black:b]" + highlight + "[white:black]"
				log = strings.Replace(log, "LOG", highlight, 1)
			case "DGB":
				highlight = "[purple:black:b]" + highlight + "[white:black]"
				log = strings.Replace(log, "DGB", highlight, 1)
			case "FTL":
				highlight = "[red:black:b]" + highlight + "[white:black]"
				log = strings.Replace(log, "FTL", highlight, 1)
			}
			fmt.Fprintln(v, log)
		}
	}
}

func saveLog(log string) {
	logStream <- log
}

type Logger struct {
	packageName string
}

func (l *Logger) Init(packageName string) error {
	l.packageName = packageName
	return nil
}

// Log func: print info msg and log it
func (l *Logger) Log(msg string) {
	logStr := fmt.Sprintf(" %v [T:%d] [M: %v ]", l.packageName, time.Now().Unix(), msg)
	go saveLog("LOG" + logStr)
}

func (l *Logger) Debug(msg string) {
	debugStr := fmt.Sprintf(" %v [T:%d] [M: %v ]", l.packageName, time.Now().Unix(), msg)
	go saveLog("DGB" + debugStr)
}

// Fail func: print info msg and log it
func (l *Logger) Fatal(err error) {
	fatalStr := fmt.Sprintf(" %v [T:%d] [M: %v ]", l.packageName, time.Now().Unix(), err.Error())
	saveLog("FTL" + fatalStr)
	os.Exit(1)
}
