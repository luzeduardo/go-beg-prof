package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/labstack/gommon/log"
)

type Configuration struct {
	Port            string `json:"port"`
	Defaultlanguage string `json:"default_language"`
	LegacyEndpoint  string `json:"legacy_endpoint"`
	DatabaseType    string `json:"database_type"`
	DatabaseURL     string `json:"database_url"`
}

var defaultConfiguration = Configuration{
	Port:            ":8080",
	Defaultlanguage: "english",
}

func (c *Configuration) LoadEnv() {
	if lang := os.Getenv("DEFAULT_LANGUAGE"); lang != "" {
		c.Defaultlanguage = lang
	}

	if port := os.Getenv("PORT"); port != "" {
		c.Port = port
	}
}

func (c *Configuration) ParsePort() {
	if c.Port[0] != ':' {
		c.Port = ":" + c.Port
	}

	if _, err := strconv.Atoi(string(c.Port[1])); err != nil {
		fmt.Printf("invalid port %s\n", c.Port)
		c.Port = defaultConfiguration.Port
	}
}

func (c *Configuration) LoadFromJSON(path string) error {
	log.Printf("loading from file %s\n", path)
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return errors.New("unable to load configuration")
	}
	if err := json.Unmarshal(b, c); err != nil {
		log.Printf("unable to parse file %s\n", err.Error())
		return errors.New("unable to load configuration")
	}

	if c.Port == "" {
		log.Printf("empty port, reverting to default")
		c.Port = defaultConfiguration.Port
	}

	if c.Defaultlanguage == "" {
		log.Printf("empty language, reverting to default")
		c.Defaultlanguage = defaultConfiguration.Defaultlanguage
	}
	return nil
}
