package config

import (
	"github.com/spf13/viper"
)

type tomlConfig struct {
	Database database
}

type database struct {
	Host     string `mapstructure:"host" toml:"host" yaml:"host" json:"host"`
	Port     int    `mapstructure:"port" toml:"port" yaml:"port" json:"port"`
	User     string `mapstructure:"user" toml:"user" yaml:"user" json:"user"`
	Password string `mapstructure:"password" toml:"password" yaml:"password" json:"password"`
	Database string `mapstructure:"database" toml:"database" yaml:"database" json:"database"`
}

var Cfg *tomlConfig

func init() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AddConfigPath("..")
	viper.SetConfigType("toml")
	viper.SetEnvPrefix("PICCOLLECTOR")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	Cfg = &tomlConfig{}
	err = viper.Unmarshal(Cfg)
	if err != nil {
		panic(err)
	}
}
