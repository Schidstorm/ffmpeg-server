package controller

import (
	"ffmpeg-server/view"
	"net/http"
)

type Static struct {
	ViewName string
}

func CreateStaticController(filename string) *Static {
	return &Static{ViewName: filename}
}

func (c *Static) Handle(res http.ResponseWriter, req *http.Request) error {
	_, _ = res.Write(view.Static(c.ViewName))
	return nil
}