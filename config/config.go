package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type (
	Config struct {
		HTTP     `yaml:"http"`
		Postgres `yaml:"postgres"`
	}

	HTTP struct {
		Port string `env-required:"true" yaml:"port"`
	}

	Postgres struct {
		Host     string `env-required:"true" yaml:"host"`
		Port     string `env-required:"true" yaml:"port"`
		Username string `env-required:"true" yaml:"username"`
		Password string `env-required:"true" env:"PG_PASSWORD"`
		DBName   string `env-required:"true" yaml:"dbname"`
		SSLMode  string `env-required:"true" yaml:"sslmode"`
	}
)

func NewConfig() (*Config, error) {
	cfg := &Config{}

	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	if err := cleanenv.ReadConfig("./config/config.yml", cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
