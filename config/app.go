package config

type AppConfig struct {
	Name          string `yaml:"name"`
	Version       string `yaml:"version"`
	Environment   string `yaml:"-"`
	DebugMode     bool   `yaml:"debug"`
	ServerAppHost string `yaml:"host"`
	ServerAppPort string `yaml:"port"`
}
