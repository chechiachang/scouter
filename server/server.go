package server

import (
	"log"
	"net"
	"net/http"

	"github.com/chechiachang/scouter"
	"github.com/linkernetworks/logger"
)

// App is the structure to set config & service provider of APP
type Apiserver struct {
	Config          scouter.Config
	ServiceProvider *scouter.Container
}

// LoadConfig consumes a string of path to the json config file and read config file into Config.
func (a *Apiserver) LoadConfig(configPath string) *Apiserver {
	if configPath == "" {
		log.Fatal("-config option is required.")
	}

	a.Config = scouter.MustRead(configPath)
	return a
}

// Start consumes two strings, host and port, invoke service initilization and serve on desired host:port
func (a *Apiserver) Start(host, port string) error {

	a.InitilizeService()

	bind := net.JoinHostPort(host, port)

	return http.ListenAndServe(bind, a.AppRoute())
}

// InitilizeService weavering services with global variables inside server package
func (a *Apiserver) InitilizeService() {
	logger.Setup(a.Config.Logger)

	a.ServiceProvider = scouter.New(a.Config)
}
