package view

import "github.com/gobuffalo/packr"

var box packr.Box

func LoadTemplates() {
	box = packr.NewBox("../templates")
}

func FindString(name string) (string, error) {
	return box.FindString(name)
}

func Find(name string) ([]byte, error) {
	return box.Find(name)
}
