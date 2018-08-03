package server

import (
	"github.com/chechiachang/scouter"
	"github.com/emicklei/go-restful"
	"github.com/gorilla/mux"
	"github.com/linkernetworks/logger"
)

// AppRoute will add router
func (a *Apiserver) AppRoute() *mux.Router {
	router := mux.NewRouter()

	container := restful.NewContainer()

	container.Filter(globalLogging)

	container.Add(newVersionService(a.ServiceProvider))

	router.PathPrefix("/v1/").Handler(container)
	return router
}

func globalLogging(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
	logger.Infof("%s %s", req.Request.Method, req.Request.URL)
	chain.ProcessFilter(req, resp)
}

func newVersionService(sp *scouter.Container) *restful.WebService {
	webService := new(restful.WebService)
	webService.Path("/v1/version").Consumes(restful.MIME_JSON, restful.MIME_JSON).Produces(restful.MIME_JSON, restful.MIME_JSON)
	//  webService.Filter(validateTokenMiddleware)
	webService.Route(webService.GET("/").To(RESTfulServiceHandler(sp, versionHandler)))
	return webService
}

// RESTfulContextHandler is the interface for restfuul handler(restful.Request,restful.Response)
type RESTfulContextHandler func(*Context)

// RESTfulServiceHandler is the wrapper to combine the RESTfulContextHandler with our scouter object
func RESTfulServiceHandler(sp *scouter.Container, handler RESTfulContextHandler) restful.RouteFunction {
	return func(req *restful.Request, resp *restful.Response) {
		ctx := Context{
			ServiceProvider: sp,
			Request:         req,
			Response:        resp,
		}
		handler(&ctx)
	}
}
