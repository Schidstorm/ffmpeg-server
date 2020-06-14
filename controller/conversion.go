package controller

import (
	"errors"
	"ffmpeg-server/model"
	"ffmpeg-server/view"
	"net/http"
)

type Conversion struct {

}

func (c *Conversion) Handle(res http.ResponseWriter, req *http.Request) error {
	switch req.Method {
	case "GET":
		for _, task := range model.GetAllTasks().Tasks {
			task.Update()
		}
		_, _ = res.Write(view.Parameterized("conversion.html", model.GetAllTasks()))
		return nil
	case "POST":
		data := model.CreateTask{}
		err := req.ParseForm()
		if err != nil {
			return err
		}
		data.TargetFileName = req.Form.Get("targetFileName")
		data.SourceFilePath = req.Form.Get("sourceFilePath")
		_ = model.NewFfmpegTask(data.SourceFilePath, data.TargetFileName)

		return NewRedirect(req.URL.String()).Handle(res, req)
	default:
		return NewError(errors.New("unknown method")).Handle(res, req)
	}
}