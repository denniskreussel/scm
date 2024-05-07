package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	LogLevel string
	LogTrace bool
}

func ParseLocalCfgFile(cfgPath string) (cfg *Config, err error) {
	cfg = &Config{}
	var b []byte
	b, err = os.ReadFile(cfgPath)
	if err != nil {
		err = fmt.Errorf("could not read config file: %w", err)
		return
	}
	if err = yaml.Unmarshal(b, cfg); err != nil {
		err = fmt.Errorf("could not unmarshal config file: %w", err)
		return
	}
	return
}
