package lib

import "os/user"

type Settings struct {
	ConversionDestinationDirectory string
	UploadDestinationDirectory string
	HttpListenEndpoint string
}

var settings *Settings = initSettings()

func initSettings() *Settings {
	currentUser, err := user.Current()
	if err != nil {
		panic(err)
	}

	return &Settings{
		ConversionDestinationDirectory: currentUser.HomeDir,
		UploadDestinationDirectory: currentUser.HomeDir,
		HttpListenEndpoint: "0.0.0.0:8080",
	}
}

func GetSettings() *Settings {
	return settings
}
