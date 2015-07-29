package main

import (
	"github.com/zshift/goplay/controllers"
	"log"
	"net/http"

	"github.com/emicklei/go-restful"
)

func main() {
	ws := new(restful.WebService)
	registerRoutes(ws)
	restful.Add(ws)

	log.Println("Starting server")
	http.ListenAndServe(":8080", nil)
}

func registerRoutes(ws *restful.WebService) {
	log.Println("Registering routes")
	healthController := controllers.NewHealthController()
	healthController.RegisterRoutes(ws)
}
