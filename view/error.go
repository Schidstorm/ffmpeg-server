package view

type ErrorData struct {
	Message string
}

func Error(err error) []byte {
	return Parameterized("error.html", ErrorData{Message: err.Error()})
}
