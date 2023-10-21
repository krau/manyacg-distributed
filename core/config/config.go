package config

import (
	"github.com/spf13/viper"
)

type tomlConfig struct {
	App       appConfig       `mapstructure:"app" toml:"app" yaml:"app" json:"app"`
	Database  database        `mapstructure:"database" toml:"database" yaml:"database" json:"database"`
	Log       logConfig       `mapstructure:"log" toml:"log" yaml:"log" json:"log"`
	Messenger messengerConfig `mapstructure:"messenger" toml:"messenger" yaml:"messenger" json:"messenger"`
}

type appConfig struct {
	Debug   bool   `mapstructure:"debug" toml:"debug" yaml:"debug" json:"debug"`
	Address string `mapstructure:"address" toml:"address" yaml:"address" json:"address"`
	CertFile string `mapstructure:"cert" toml:"cert" yaml:"cert" json:"cert"`
	KeyFile string `mapstructure:"key" toml:"key" yaml:"key" json:"key"`
	CaFile string `mapstructure:"ca" toml:"ca" yaml:"ca" json:"ca"`
}

type logConfig struct {
	Level     string `mapstructure:"level" toml:"level" yaml:"level" json:"level"`
	FilePath  string `mapstructure:"file_path" toml:"file_path" yaml:"file_path" json:"file_path"`
	BackupNum uint   `mapstructure:"backup_num" toml:"backup_num" yaml:"backup_num" json:"backup_num"`
}

type messengerConfig struct {
	Azure azureConfig `mapstructure:"azure" toml:"azure" yaml:"azure" json:"azure"`
}

type azureConfig struct {
	BusConnectionString string `mapstructure:"bus_connection_string" toml:"bus_connection_string" yaml:"bus_connection_string" json:"bus_connection_string"`
	SubTopic            string `mapstructure:"sub_topic" toml:"sub_topic" yaml:"sub_topic" json:"sub_topic"`
	Subscription        string `mapstructure:"subscription" toml:"subscription" yaml:"subscription" json:"subscription"`
	PubTopic            string `mapstructure:"pub_topic" toml:"pub_topic" yaml:"pub_topic" json:"pub_topic"`
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
	viper.SetDefault("app.address", "0.0.0.0:39010")
	viper.SetDefault("log.level", "info")
	viper.SetDefault("log.file_path", "./logs/storage.log")
	viper.SetDefault("log.backup_num", 7)
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
