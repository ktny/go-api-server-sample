package config

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

var Config Conf

type Environment struct {
	Dev Conf `yaml: "dev"`
}

type Conf struct {
	Database Database `yaml:"db"`
}

type Database struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
}

func SetEnvironment(env string) {
	buf, err := ioutil.ReadFile("./config/env.yml")
	if err != nil {
		panic(err)
	}

	var environment Environment

	err = yaml.Unmarshal(buf, &environment)
	if err != nil {
		panic(err)
	}

	if env == "dev" {
		Config = environment.Dev
	}
}
