package view

import "github.com/gobuffalo/packr/v2"

var box *packr.Box

func LoadTemplates() {

	box = packr.New("templates", "../templates")
}

func FindString(name string) (string, error) {
	return box.FindString(name)
}

func Find(name string) ([]byte, error) {
	return box.Find(name)
}
