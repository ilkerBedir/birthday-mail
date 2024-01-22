package config

import (
	"log"
	"os"
	"sync"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Database struct {
		URL string `yaml:"URL"`
	} `yaml:"database"`
}

var config *Config
var lock = &sync.Mutex{}

func GetConfig() *Config {
	if config == nil {
		lock.Lock()
		defer lock.Unlock()
		log.Println("Config yükleniyor")
		f, err := os.Open("conf.yml")
		if err != nil {
			log.Fatalln("Conf yüklenmedi : ", err)
		}
		defer f.Close()
		decoder := yaml.NewDecoder(f)
		err = decoder.Decode(&config)
		if err != nil {
			log.Fatalln("Conf yüklenmedi : ", err)
		}
		return config
	}

	return config
}
