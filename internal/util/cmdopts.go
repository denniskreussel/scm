package util

import (
	"os"
	"strconv"
	"strings"
)

type CmdOptions struct {
	ShowVer  bool
	LogLevel string // "debug", "info", "warn", "error", "dpanic", "panic", "fatal"
	LogPaths string // comma-separated paths. "stdout" means the console stdout

	// HTTPHost to bind to. If empty, outbound ip of machine
	// is automatically determined and used.
	HTTPHost string
	HTTPPort int // 0 means a randomly chosen port.

	PushGatewayAddrs string
	PushInterval     int
	LocalCfgFile     string
}

func EnvStringVar(value *string, key string) {
	realKey := strings.ReplaceAll(strings.ToUpper(key), "-", "_")
	val, found := os.LookupEnv(realKey)
	if found {
		*value = val
	}
}

func EnvIntVar(value *int, key string) {
	realKey := strings.ReplaceAll(strings.ToUpper(key), "-", "_")
	val, found := os.LookupEnv(realKey)
	if found {
		valInt, err := strconv.Atoi(val)
		if err == nil {
			*value = valInt
		}
	}
}

func EnvBoolVar(value *bool, key string) {
	realKey := strings.ReplaceAll(strings.ToUpper(key), "-", "_")
	if _, found := os.LookupEnv(realKey); found {
		*value = true
	}
}
