package initialize

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"server/config"
)

func ConfigInit(path string) *config.Config {
	var c config.Config
	configFile, err := ioutil.ReadFile(path + "config.yaml")
	if err != nil {
		panic(err)
	}
	err2 := yaml.Unmarshal(configFile, &c)
	if err2 != nil {
		panic(err)
	}
	return &c
}
