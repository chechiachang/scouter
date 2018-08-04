package server

import (
	"github.com/linkernetworks/logger"
)

type FaceLandmarks struct {
}

type RecognitionResponse struct {
	ID            int64  `json:"id"`
	Login         string `json:"login"`
	Publicrepos   int    `json:"publicrepos"`
	Followers     int    `json:"followers"`
	Contributions int    `json:"contributions"`
}

func recognitionHandler(ctx *Context) {
	_, _, resp := ctx.ServiceProvider, ctx.Request, ctx.Response

	var faceLandmarks FaceLandmarks
	if err := ctx.Request.ReadEntity(&faceLandmarks); err != nil {
		logger.Error(err)
	}
	resp.WriteEntity(RecognitionResponse{
		ID:            int64(0),
		Login:         "chechiachang",
		Publicrepos:   50,
		Followers:     60,
		Contributions: 100,
	})
}
