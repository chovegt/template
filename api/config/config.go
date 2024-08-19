package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type AppConfig struct {
	Log struct {
		Lvl string `yaml:"lvl"`
	} `yaml:"log"`
}

func Get(basePath, scope string) (AppConfig, error) {
	dataConfig := AppConfig{}
	file := basePath + scope + "/config.yml"

	yamlFile, err := os.ReadFile(file)
	if err != nil {
		return dataConfig, err
	}

	err = yaml.Unmarshal(yamlFile, &dataConfig)
	if err != nil {
		return dataConfig, err
	}

	return dataConfig, err
}
