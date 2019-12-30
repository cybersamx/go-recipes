package pkg

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type Settings struct {
	Port int    `yaml:"port"`
	DSN  string `yaml:"dsn"`
}

func NewSettings() *Settings {
	const filename = "env.yaml"
	var defaultSettings = Settings{
		Port: 3000,
		DSN:  ":memory",
	}

	info, err := os.Stat(filename)
	if err != nil || info.IsDir() {
		return &defaultSettings
	}

	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return &defaultSettings
	}

	settings := Settings{}
	err = yaml.Unmarshal(content, &settings)
	if err != nil {
		return &defaultSettings
	}

	return &settings
}
