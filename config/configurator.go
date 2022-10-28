package config

import (
	"github.com/spf13/viper"
	"log"
)

type ServerConfig struct {
	Host string `yaml:"server.host"`
	Port string `yaml:"server.port"`
	API  struct {
		UnixSocket string `yaml:"server.api.unix_socket"`
	}
}
type DatabaseConfig struct {
	Host     string `yaml:"db.host"`
	User     string `yaml:"db.user"`
	Password string `yaml:"db.password"`
	Name     string `yaml:"db.name"`
	Port     string `yaml:"db.port"`
	Ssl      string `yaml:"db.ssl"`
	Timezone string `yaml:"db.timezone"`
	Network  string `yaml:"db.network"`
	Charset  string `yaml:"db.charset"`
	Migrate  bool   `yaml:"db.migrate"`
}
type Config struct {
	Environment string         `yaml:"environment"`
	Server      ServerConfig   `mapstructure:"server"`
	Database    DatabaseConfig `mapstructure:"database"`
}

type Configurator interface {
	GetConfiguration() *Config
}

func NewConfigurator(fileName string, filePath string) *Config {
	return NewYamlConfigurator(fileName, filePath).GetConfiguration()
}

type YamlConfigurator struct {
	config Config
}

func NewYamlConfigurator(fileName string, filePath string) Configurator {
	var configuration Config

	v := viper.New()
	v.SetConfigType("yaml")
	v.SetConfigName(fileName)
	v.AddConfigPath(filePath)

	err := v.ReadInConfig()
	if err != nil {
		log.Panic("Error on reading config file.", err)
	}

	err = v.Unmarshal(&configuration)
	if err != nil {
		log.Panic("Error on parsing config file.", err)
	}

	return &YamlConfigurator{config: configuration}
}

func (yc *YamlConfigurator) GetConfiguration() *Config {
	return &yc.config
}
