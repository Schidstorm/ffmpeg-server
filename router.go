package main

import (
	"encoding/json"
	"ffmpeg-server/controller"
	"github.com/sirupsen/logrus"
	"net/http"
)

type RequestHandler interface {
	Handle(res http.ResponseWriter, req *http.Request) error
}

type Route struct {
	Path       string
	Controller RequestHandler
}

func routes() []Route {
	return []Route{
		{Path: "/", Controller: controller.CreateStaticController("index.html")},
		{Path: "/favicon.ico", Controller: controller.CreateStaticController("favicon.ico")},
		{Path: "/dropzone.min.js", Controller: controller.CreateStaticController("dropzone.min.js")},
		{Path: "/dropzone.min.css", Controller: controller.CreateStaticController("dropzone.min.css")},
		{Path: "/conversion", Controller: new(controller.Conversion)},
		{Path: "/upload", Controller: new(controller.Upload)},
		{Path: "/tasks", Controller: new(controller.Tasks)},
	}
}

type ErrorResponse struct {
	Error string
}

func applyRoutes() {
	for _, r := range routes() {
		route := r
		http.HandleFunc(route.Path, func(writer http.ResponseWriter, request *http.Request) {
			err := route.Controller.Handle(writer, request)
			if err != nil {
				logrus.Error(err)
				writer.Header().Add("Content-Type", "application/json; charset=utf-8")
				writer.WriteHeader(500)
				data, _ := json.Marshal(ErrorResponse{Error: err.Error()})
				_, _ = writer.Write(data)
			}
		})
	}
}