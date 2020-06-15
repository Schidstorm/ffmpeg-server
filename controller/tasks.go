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
	data, err := json.Marshal(model.GetAllTasks())
	if err != nil {
		return err
	}
	res.Header().Add("Content-Type", "application/json")
	res.WriteHeader(200)
	_, err = res.Write(data)
	if err != nil {
		return err
	}
	return nil
}