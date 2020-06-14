package controller

import (
	"errors"
	"ffmpeg-server/lib"
	"ffmpeg-server/model"
	"ffmpeg-server/view"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
	"path"
)

type Upload struct {
	model model.Upload
}

func (c *Upload) Handle(res http.ResponseWriter, req *http.Request) error {
	switch req.Method {
	case "GET":
		_, _ = res.Write(view.Static("upload.html"))
	case "POST":
		reader, err := req.MultipartReader()
		if err != nil {
			return errors.New("no multipart data detected")
		}
		part, err := reader.NextPart()
		if err != nil {
			return errors.New("no uploaded file found")
		}

		uploadFilePath := path.Join(lib.GetSettings().UploadDestinationDirectory, part.FileName())
		uploadFile, err := os.OpenFile(uploadFilePath, os.O_CREATE | os.O_RDWR, 0660)
		if err != nil {
			return err
		}
		writtenBytes, err := io.Copy(uploadFile, part)
		if err != nil {
			return err
		}
		logrus.Println(writtenBytes, "bytes written")
		_ = uploadFile.Close()
		_ = part.Close()

		return NewRedirect(req.URL.String()).Handle(res, req)
	}
	return nil
}