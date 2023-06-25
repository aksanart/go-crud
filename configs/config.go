package configs

import (
	"errors"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	App   AppConfig
	Mysql MysqlConfig
}

type AppConfig struct {
	Name string
	Port int
	Mode string
}

type MysqlConfig struct {
	Host     string
	Username string
	Password string
	Database string
	Port     int
}

// LoadConfig file from given path
func LoadConfig(filename string) (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigName(filename)
	v.AddConfigPath(".")
	v.AutomaticEnv()
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Println("err", err)
			return nil, errors.New("config file not found")
		}
		return nil, err
	}
	return v, nil
}

// ParseConfig file
func ParseConfig(v *viper.Viper) (*Config, error) {
	var c Config
	err := v.Unmarshal(&c)
	if err != nil {
		log.Printf("unable to decode into struct, %v", err)
		return nil, err
	}

	return &c, nil
}

func GetConfig(configPath string) (*Config, error) {
	cfgFile, err := LoadConfig(configPath)
	if err != nil {
		return nil, err
	}

	cfg, err := ParseConfig(cfgFile)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

func GetConfigPath(configPath string) string {
	viper.ReadInConfig()
	return "./configs/env"
}
