package data

import (
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

var ConfigCache *Config

type Config struct {
	DataDirectory string        `yaml:"data_directory"`
	MinioSettings MinioSettings `yaml:"minio_settings"`
}

type MinioSettings struct {
	Buckets   []string `yaml:"buckets"`
	Endpoint  string   `yaml:"endpoint"`
	AccessKey string   `yaml:"access_key"`
	SecretKey string   `yaml:"secret_key"`
	UseSSL    bool     `yaml:"use_ssl"`
}

func InitConfig() {
	stop := make(chan os.Signal, 1)

	if ConfigCache == nil {
		yamlFile, err := ioutil.ReadFile("config.yml")
		if err != nil {
			log.Panic("Failed to read config.yml.")
			<-stop
		}
		ConfigCache = &Config{}
		err = yaml.Unmarshal(yamlFile, ConfigCache)
		if err != nil {
			log.Panic("Failed to parse config.yml.")
			<-stop
		}
	}
}
