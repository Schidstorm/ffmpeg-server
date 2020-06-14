package lib

import "os/user"

type Settings struct {
	ConversionDestinationDirectory string
	UploadDestinationDirectory string
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
	}
}

func GetSettings() *Settings {
	return settings
}
