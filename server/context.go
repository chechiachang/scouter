package server

import (
	"github.com/chechiachang/scouter/serviceprovider"
	"github.com/emicklei/go-restful"
)

// Context is the struct to combine the restful message with our own serviceProvider
type Context struct {
	ServiceProvider *serviceprovider.Container
	Request         *restful.Request
	Response        *restful.Response
}
