package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

var (
	Global  GlobalConfig
	FileUrl string = "./config/env.yml"
)

type (
	ServerConfig struct {
		Port string `yaml:"port"`
	}

	DatabaseConfig struct {
		Mode        string `yaml:"mode"`
		Datasource  string `yaml:"datasource"`
		MaxIdleConn int    `yaml:"max_idle_conns"`
		MaxOpenConn int    `yaml:"max_open_conns"`
	}

	GlobalConfig struct {
		Server   ServerConfig   `yaml:"server"`
		Database DatabaseConfig `yaml:"database"`
	}
)

func Boot() (globalConfig GlobalConfig, err error) {
	data, err := os.ReadFile(FileUrl)
	if err != nil {
		log.Printf("%v", err)
	}

	yaml.Unmarshal(data, &Global)

	return Global, nil
}
