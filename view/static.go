package view

import "github.com/sirupsen/logrus"

func Static(viewName string) []byte {
	html, err := Find(viewName)
	if err != nil {
		logrus.Error(err, viewName)
	}
	return html
}
