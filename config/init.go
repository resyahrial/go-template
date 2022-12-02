package config

import (
	"fmt"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

const (
	cfgPath = "./files/etc/app_config/config.%s.yml"
)

type (
	Config struct {
		App      AppConfig      `yaml:"app"`
		Database DatabaseConfig `yaml:"database"`
	}
)

var (
	GlobalConfig Config
)

func InitConfig(env string) {
	if err := cleanenv.ReadConfig(getConfigPath(env), &GlobalConfig); err != nil {
		log.Fatalf("Failed to read config: %v", err)
		os.Exit(2)
	}
}

func getConfigPath(env string) string {
	return fmt.Sprintf(cfgPath, env)
}
