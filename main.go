package main

import (
	"fmt"
	"log"

	"github.com/radhianamri/go-libs/config"
)

// Config represents a generic configuration structure
type Config struct {
	// Add fields based on your configuration needs
	Name     string `json:"name" yaml:"name" toml:"name" ini:"name"`
	Age      int    `json:"age" yaml:"age" toml:"age" ini:"age"`
	Email    string `json:"email" yaml:"email" toml:"email" ini:"email"`
	Database struct {
		Host     string `json:"host" yaml:"host" toml:"host" ini:"host"`
		Port     int    `json:"port" yaml:"port" toml:"port" ini:"port"`
		Username string `json:"username" yaml:"username" toml:"username" ini:"username"`
		Password string `json:"password" yaml:"password" toml:"password" ini:"password"`
	} `json:"database" yaml:"database" toml:"database" ini:"database"`
}

func main() {
	// Example usage
	var cfg Config

	// JSON
	err := config.LoadJson("data/example.json", &cfg)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Deserialized JSON config: %+v\n", cfg)

	// YAML
	err = config.LoadYaml("data/example.yaml", &cfg)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Deserialized YAML config: %+v\n", cfg)

	// TOML
	err = config.LoadToml("data/example.toml", &cfg)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Deserialized TOML config: %+v\n", cfg)

	// INI
	err = config.LoadIni("data/example.ini", &cfg)
	if err != nil {
		log.Fatal("asdsadas", err.Error())
	}
	fmt.Printf("Deserialized INI config: %+v\n", cfg)
}
