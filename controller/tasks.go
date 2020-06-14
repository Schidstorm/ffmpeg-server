package controller

import (
	"encoding/json"
	"ffmpeg-server/model"
	"net/http"
)

type Tasks struct {

}

func (t *Tasks) Handle(res http.ResponseWriter, req *http.Request) error {
	for _, t := range model.GetAllTasks().Tasks {
		t.Update()
	}
	data, _ := json.Marshal(model.GetAllTasks())
	res.Header().Add("Content-Type", "application/json")
	res.WriteHeader(200)
	_, _ = res.Write(data)
	return nil
}