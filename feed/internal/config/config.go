package config

import (
	"fmt"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env string `yaml:"env" env-required:"true"`
	HTTPServer `yaml:"http_server" env-required:"true"`
	DB `yaml:"db" env-required:"true"`
}

type HTTPServer struct {
	Host string `yaml:"host" env-required:"true"`
	Port string `yaml:"port" env-required:"true"`
	Timeout time.Duration `yaml:"timeout" env-required:"true"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-required:"true"`
}

type DB struct {
	Username string `yaml:"username" env-required:"true"`
	Host string `yaml:"host" env-required:"true"`
	Port string `yaml:"port" env-required:"true"`
	DBName string `yaml:"db_name" env-required:"true"`
	SSLMode string `yaml:"sslmode" env-required:"true"`
}

func MustLoad() *Config {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		panic("CONFIG_PATH is not set")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		panic(fmt.Sprintf("Config is not exist on path: %s", configPath))
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		panic(fmt.Sprintf("Cannot read config: %s", err))
	}

	return &cfg
}