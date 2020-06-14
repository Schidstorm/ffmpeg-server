package main

import (
	"ffmpeg-server/view"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
)

func main() {
	view.LoadTemplates()
	applyRoutes()
	err := http.ListenAndServe("0.0.0.0:8080", nil)
	if err != nil {
		logrus.Error(err)
		os.Exit(1)
	}


}
