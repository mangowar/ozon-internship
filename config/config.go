package config

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	BaseURL  string   `yaml:"base_url"`
	Host     string   `yaml:"host"`
	Port     int      `yaml:"port"`
	DB       DBConfig `yaml:"db"`
	FileName string   `yaml:"file_name"`
}

func (c Config) ListenAddr() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}

type DBConfig struct {
	DSN      string `yaml:"DSN"`
	Database string `yaml:"database"`
}

func Get() Config {
	var cfg Config
	file, err := os.Open("config/config.yaml")
	if err != nil {
		log.Fatalf("No such file: %v", err)
	}
	defer file.Close()
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&cfg); err != nil {
		log.Fatal("yaml decode error")
	}
	return cfg
}
