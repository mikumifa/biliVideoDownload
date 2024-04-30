package config

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type Config struct {
	HttpConfig HttpConfig `yaml:"http"`
}

type HttpConfig struct {
	Cookie string `yaml:"cookie"`
}

var config *Config

func init() {
	// 加载配置
	err := load("config.yml")
	if err != nil {
		logrus.Info("config.yml not found", err)
		return
	}
}

func load(path string) error {
	result, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(result, &config)
}

func Get() *Config {
	return config
}
