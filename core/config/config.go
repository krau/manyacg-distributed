package config

import (
	"github.com/spf13/viper"
)

type tomlConfig struct {
	App      appConfig `mapstructure:"app" toml:"app" yaml:"app" json:"app"`
	Database database  `mapstructure:"database" toml:"database" yaml:"database" json:"database"`
}

type appConfig struct {
	Log   logConfig   `mapstructure:"log" toml:"log" yaml:"log" json:"log"`
	Debug bool        `mapstructure:"debug" toml:"debug" yaml:"debug" json:"debug"`
	Azure azureConfig `mapstructure:"azure" toml:"azure" yaml:"azure" json:"azure"`
}

type logConfig struct {
	Level     string `mapstructure:"level" toml:"level" yaml:"level" json:"level"`
	FilePath  string `mapstructure:"file_path" toml:"file_path" yaml:"file_path" json:"file_path"`
	BackupNum uint   `mapstructure:"backup_num" toml:"backup_num" yaml:"backup_num" json:"backup_num"`
}

type azureConfig struct {
	BusConnectionString string `mapstructure:"bus_connection_string" toml:"bus_connection_string" yaml:"bus_connection_string" json:"bus_connection_string"`
	Topic               string `mapstructure:"topic" toml:"topic" yaml:"topic" json:"topic"`
	Subscription        string `mapstructure:"subscription" toml:"subscription" yaml:"subscription" json:"subscription"`
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
	viper.SetDefault("app.debug", false)
	viper.SetDefault("app.log.level", "info")
	viper.SetDefault("app.log.file_path", "./logs/core.log")
	viper.SetDefault("app.log.backup_num", 7)
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
