package main

import (
	"flag"

	"github.com/chechiachang/scouter/server"
)

func main() {
	var (
		configPath string
		host       string
		port       string
	)

	flag.StringVar(&configPath, "config", "config.json", "config file path")
	flag.StringVar(&host, "host", "0.0.0.0", "hostname")
	flag.StringVar(&port, "port", "5487", "port")

	flag.Parse()

	a := server.Apiserver{}
	a.LoadConfig(configPath).Start(host, port)
}
