package controller

import (
	"encoding/json"
	"ffmpeg-server/lib"
	"io/ioutil"
	"net/http"
)

type SourceFiles struct {

}

func (c *SourceFiles) Handle(res http.ResponseWriter, req *http.Request) error {
	fileInfos, err := ioutil.ReadDir(lib.GetSettings().UploadDestinationDirectory)
	if err != nil {
		return err
	}

	files := make([]string, 0)
	for _, info := range fileInfos {
		files = append(files, info.Name())
	}

	filesJson, err := json.Marshal(files)
	if err != nil {
		return err
	}

	res.Header().Add("Content-Type", "application/json")
	res.WriteHeader(200)
	res.Write(filesJson)

	return nil
}