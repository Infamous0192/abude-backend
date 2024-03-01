package config

import (
	"fmt"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type AppConfig struct {
	Database   DatabaseConfig `yaml:"database"`
	Server     ServerConfig   `yaml:"server"`
	Jwt        JwtConfig      `yaml:"jwt"`
	ConfigFile string
}

func Load(configFile string) *AppConfig {
	cfg := AppConfig{ConfigFile: configFile}
	if err := cleanenv.ReadConfig(cfg.ConfigFile, &cfg); err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	return &cfg
}
