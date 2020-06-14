package controller

import "net/http"

type Redirect struct {
	Url string
}

func (r *Redirect) Handle(res http.ResponseWriter, req *http.Request) error {
	res.Header().Add("Location", r.Url)
	res.WriteHeader(302)
	return nil
}

func NewRedirect(url string) *Redirect {
	return &Redirect{Url: url}
}