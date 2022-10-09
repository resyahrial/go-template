package config

import (
	"fmt"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/sirupsen/logrus"
)

const (
	cfgPath = "./files/etc/app_config/config.%s.yaml"
)

type (
	Config struct {
		RestServer RestServerConfig `yaml:"rest_server"`
	}
)

var (
	GlobalConfig Config
)

func InitConfig(env string) {
	if err := cleanenv.ReadConfig(getConfigPath(env), &GlobalConfig); err != nil {
		logrus.Fatalf("Failed to read config: %v", err)
		os.Exit(2)
	}
}

func getConfigPath(env string) string {
	return fmt.Sprintf(cfgPath, env)
}
