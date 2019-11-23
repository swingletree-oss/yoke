package swingletree

import (
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Yoke yokeconf
}

type yokeconf struct {
	Reports []Report
}

type Report struct {
	Plugin      string
	Report      string
	ContentType string
}

func LoadConf(configFile string) (Config, error) {
	c := Config{}
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		return Config{}, err
	} else {
		yamlFile, err := ioutil.ReadFile(configFile)
		if err != nil {
			log.Fatalf("Failed to load configuration. Caused by; %v ", err)
		}
		err = yaml.Unmarshal(yamlFile, &c)
		if err != nil {
			log.Fatalf("Failed to unmarshal configuration contents. Caused by; %v ", err)
		}
	}

	return c, nil
}
