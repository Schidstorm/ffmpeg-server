package view

import (
	"ffmpeg-server/lib"
	"html/template"
)

func Parameterized(viewName string, data interface{}) []byte {
	html := template.New("page")
	htmlString, _ := FindString(viewName)
	_, _ = html.Parse(htmlString)
	memoryWriter := lib.CreateMemoryWriter()
	_ = html.Execute(memoryWriter, data)
	return memoryWriter.Buffer()
}
