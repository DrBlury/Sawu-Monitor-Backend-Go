package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Conf struct {
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Sawu     struct {
		Separator string `yaml:"separator"`
	}
	Kafka struct {
		Broker struct {
			IPAddress string `yaml:"ipaddress"`
		}
		Consumer struct {
			ConsumerGroup string   `yaml:"consumergroup"`
			Topics        []string `yaml:"topics"`
		}
	}
	Database struct {
		Server   string `yaml:"server"`
		Port     string `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Schema   string `yaml:"schema"`
	}
}

func (c *Conf) GetDefaults() *Conf {

	yamlFile, err := ioutil.ReadFile("./res/defaults.yml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}
