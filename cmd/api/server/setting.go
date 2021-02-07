package server

import (
	"github.com/kelseyhightower/envconfig"
)

type ProjectSettings struct {
	ProjectName    string `default:"agile-engine-test"`
	ProjectVersion string `default:"1.0.0"`
	UrlBase        string `envconfig:"BASE_URL" required:"true" default:"localhost"`
	Host           string `envconfig:"HOST" default:"0.0.0.0"`
	Port           string `envconfig:"PORT" default:"8080"`
}

func InitializeSettings() ProjectSettings {
	var Settings ProjectSettings
	if err := envconfig.Process("", &Settings); err != nil {
		panic(err.Error())
	}
	return Settings
}