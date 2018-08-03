package server

import (
	"github.com/chechiachang/scouter"
	"github.com/emicklei/go-restful"
)

// Context is the struct to combine the restful message with our own serviceProvider
type Context struct {
	ServiceProvider *scouter.Container
	Request         *restful.Request
	Response        *restful.Response
}
