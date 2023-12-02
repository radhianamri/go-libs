package config

import (
	"os"

	"github.com/go-ini/ini"
	"github.com/pelletier/go-toml"
	"github.com/radhianamri/go-libs/json"
	"gopkg.in/yaml.v2"
)

// Load deserializes a configuration file based on the provided format
func LoadToml(filename string, config interface{}) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	tree, err := toml.Load(string(data))
	if err != nil {
		return err
	}
	return tree.Unmarshal(config)
}

// Load deserializes a configuration file based on the provided format
func LoadYaml(filename string, config interface{}) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(data, config)
}

// Load deserializes a configuration file based on the provided format
func LoadJson(filename string, config interface{}) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, config)
}

// Load deserializes a configuration file based on the provided format
func LoadIni(filename string, config interface{}) error {
	return ini.MapTo(config, filename)
}
