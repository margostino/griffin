package config

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

func LoadCommands(configFile string) *CommandsConfiguration {
	var commandsConfiguration CommandsConfiguration
	yamlFile, err := ioutil.ReadFile(configFile)

	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, &commandsConfiguration)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return &commandsConfiguration
}
