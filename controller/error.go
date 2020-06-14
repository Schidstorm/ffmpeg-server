package controller


import (
	"ffmpeg-server/view"
	"net/http"
)

type Error struct {
	Error error
}

func NewError(err error) *Error {
	return &Error{Error: err}
}

func (c *Error) Handle(res http.ResponseWriter, req *http.Request) error {
	res.WriteHeader(500)
	res.Write(view.Error(c.Error))
	return nil
}