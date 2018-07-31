package server

import (
	"github.com/emicklei/go-restful"
	"github.com/linkernetworks/logger"
)

func globalLogging(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
	logger.Infof("%s %s", req.Request.Method, req.Request.URL)
	chain.ProcessFilter(req, resp)
}
