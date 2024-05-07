package main

import (
	"fmt"
	"github.com/denniskreussel/scm/config"
)

func init() {
	// This will be executed before main
}

func main() {
	// This is the main entry point
	cfg, err := config.ParseLocalCfgFile("dev/scm.config.yaml")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Config: %+v\n", cfg)
}
