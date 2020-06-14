package main

import (
	"ffmpeg-server/lib"
	"ffmpeg-server/view"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
)

func main() {
	view.LoadTemplates()
	applyRoutes()
	logrus.Println("Listening on ", "http://"+lib.GetSettings().HttpListenEndpoint)
	err := http.ListenAndServe(lib.GetSettings().HttpListenEndpoint, nil)
	if err != nil {
		logrus.Error(err)
		os.Exit(1)
	}


}
