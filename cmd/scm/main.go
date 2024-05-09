package main

import (
	"flag"
	"fmt"
	"github.com/denniskreussel/scm/config"
	"github.com/denniskreussel/scm/internal/util"
	"os"
	"strings"
)

var (
	version = "None"
	commit  = "None"
	date    = "None"
	builtBy = "None"

	cmdOps util.CmdOptions
)

func init() {
	initCmdOptions()
	logPaths := strings.Split(cmdOps.LogPaths, ",")
	util.InitLogger(logPaths)
	util.SetLogLevel(cmdOps.LogLevel)
	util.Logger.Info(getVersion())
	if cmdOps.ShowVer {
		os.Exit(0)
	}
}

func main() {
	// This is the main entry point
	cfg, err := config.ParseLocalCfgFile(cmdOps.LocalCfgFile)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Config: %+v\n", cfg)
}

func initCmdOptions() {
	cmdOps = util.CmdOptions{
		LogLevel:     "info",
		LogPaths:     "stdout,scm.log",
		PushInterval: 10,
		LocalCfgFile: "/etc/scm/scm.config.yaml",
	}

	// 2. Replace options with the corresponding env variable if present.
	util.EnvBoolVar(&cmdOps.ShowVer, "v")
	util.EnvStringVar(&cmdOps.LogLevel, "log-level")
	util.EnvStringVar(&cmdOps.LogPaths, "log-paths")
	util.EnvIntVar(&cmdOps.HTTPPort, "http-port")
	util.EnvStringVar(&cmdOps.HTTPHost, "http-host")
	util.EnvStringVar(&cmdOps.PushGatewayAddrs, "metric-push-gateway-addrs")
	util.EnvIntVar(&cmdOps.PushInterval, "push-interval")
	util.EnvStringVar(&cmdOps.LocalCfgFile, "local-cfg-file")

	// 3. Replace options with the corresponding CLI parameter if present.
	flag.BoolVar(&cmdOps.ShowVer, "v", cmdOps.ShowVer, "show build version and quit")
	flag.StringVar(&cmdOps.LogLevel, "log-level", cmdOps.LogLevel, "one of debug, info, warn, error, dpanic, panic, fatal")
	flag.StringVar(&cmdOps.LogPaths, "log-paths", cmdOps.LogPaths, "a list of comma-separated log file path. stdout means the console stdout")
	flag.IntVar(&cmdOps.HTTPPort, "http-port", cmdOps.HTTPPort, "http listen port")
	flag.StringVar(&cmdOps.HTTPHost, "http-host", cmdOps.HTTPHost, "http host to bind to")
	flag.StringVar(&cmdOps.PushGatewayAddrs, "metric-push-gateway-addrs", cmdOps.PushGatewayAddrs, "a list of comma-separated prometheus push gatway address")
	flag.IntVar(&cmdOps.PushInterval, "push-interval", cmdOps.PushInterval, "push interval in seconds")
	flag.StringVar(&cmdOps.LocalCfgFile, "local-cfg-file", cmdOps.LocalCfgFile, "local config file")

	flag.Parse()
}

func getVersion() string {
	return fmt.Sprintf("version %s, commit %s, date %s, builtBy %s, pid %v", version, commit, date, builtBy, os.Getpid())
}
