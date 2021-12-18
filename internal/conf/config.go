package conf

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Elastic Elastic `json:"elastic" yaml:"elastic"`
	Server  Server  `json:"server" yaml:"server"`
}

type Elastic struct {
	Addr string `json:"addr" yaml:"addr"`
}

type Server struct {
	Addr string `json:"addr" yaml:"addr"`
}

func LoadAndInit(conf string) *Config {
	file, err := ioutil.ReadFile(conf)
	if err != nil {
		log.Fatal("read yaml file failed")
	}

	sconf := &Config{}
	err = yaml.UnmarshalStrict(file, &sconf)
	if err != nil {
		log.Fatal("yaml unmarshal failed")
	}
	return sconf
}
