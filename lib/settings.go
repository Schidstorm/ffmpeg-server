package lib

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os/user"
)

type Settings struct {
	ConversionDestinationDirectory string
	UploadDestinationDirectory     string
	HttpListenEndpoint             string
}

var settings *Settings = initSettings()

func initSettings() *Settings {
	currentUser, err := user.Current()
	if err != nil {
		panic(err)
	}

	loadConfigFile(currentUser.HomeDir)

	return &Settings{
		ConversionDestinationDirectory: viper.GetString("ConversionDestinationDirectory"),
		UploadDestinationDirectory:     viper.GetString("UploadDestinationDirectory"),
		HttpListenEndpoint:             viper.GetString("HttpListenEndpoint"),
	}
}

func loadConfigFile(homePath string) {
	viper.SetDefault("ConversionDestinationDirectory", homePath)
	viper.SetDefault("UploadDestinationDirectory", homePath)
	viper.SetDefault("HttpListenEndpoint", "0.0.0.0:8080")

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("/etc/ffmpeg-server/")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return
		}

		logrus.Errorln(err)
	}
}

func GetSettings() *Settings {
	return settings
}
