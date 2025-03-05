package main

import (
	"flag"
	"fmt"
	"github.com/dek0valev/niwa/internal/config"
	"github.com/dek0valev/niwa/pkg/must"
)

func main() {
	fmt.Println("庭 | Niwa")

	var configPath string

	flag.StringVar(&configPath, "config", "", "path to config file")
	flag.Parse()

	cfg := must.Must(config.NewConfig(configPath))

	fmt.Println("Env:", cfg.Env)
}
