package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
)

type Config struct {
	Env string `yaml:"env" env-default:"local"`

	Content ContentConfig `yaml:"content"`
}

type ContentConfig struct {
	ArticlesDir string `yaml:"articles_dir"`
}

func NewConfig(configPath string) (*Config, error) {
	const op = "config.config.Load"

	var path string

	if configPath == "" {
		path = os.Getenv("CONFIG_PATH")
	} else {
		path = configPath
	}

	if path == "" {
		return nil, fmt.Errorf("%s: передан пустой путь до файла конфигурации", op)
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, fmt.Errorf("%s: файл конфигурации не существует: %s", op, path)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &cfg, nil
}
