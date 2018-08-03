package server

func versionHandler(ctx *Context) {
	_, _, resp := ctx.ServiceProvider, ctx.Request, ctx.Response
	resp.WriteEntity(ActionResponse{
		Error: false,
		// FIXME
		Message: "v0.0.1",
	})
}
