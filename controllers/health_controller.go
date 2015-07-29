package controllers

import (
	"io"

	"github.com/emicklei/go-restful"
)

type HealthController struct{}

func NewHealthController() *HealthController {
	return &HealthController{}
}

func (this *HealthController) RegisterRoutes(ws *restful.WebService) {
	ws.Route(ws.GET("/health").To(this.Get))
}

func (this *HealthController) Get(request *restful.Request, response *restful.Response) {
	io.WriteString(response, "ok")
}
