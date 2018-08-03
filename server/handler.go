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

// ActionResponse is the structure for Response action
type ActionResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
}
