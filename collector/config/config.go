package config

import (
	"github.com/spf13/viper"
)

type tomlConfig struct {
	App     appConfig     `mapstructure:"app" toml:"app" yaml:"app" json:"app"`
	Sources sourceConfigs `mapstructure:"sources" toml:"sources" yaml:"sources" json:"sources"`
	Sender  senderConfigs `mapstructure:"sender" toml:"sender" yaml:"sender" json:"sender"`
	Log     logConfig     `mapstructure:"log" toml:"log" yaml:"log" json:"log"`
}

type appConfig struct {
	Debug bool `mapstructure:"debug" toml:"debug" yaml:"debug" json:"debug"`
}

type logConfig struct {
	Level     string `mapstructure:"level" toml:"level" yaml:"level" json:"level"`
	FilePath  string `mapstructure:"file_path" toml:"file_path" yaml:"file_path" json:"file_path"`
	BackupNum uint   `mapstructure:"backup_num" toml:"backup_num" yaml:"backup_num" json:"backup_num"`
}

type senderConfigs struct {
	Type     string         `mapstructure:"type" toml:"type" yaml:"type" json:"type"`
	Azure    azureConfig    `mapstructure:"azure" toml:"azure" yaml:"azure" json:"azure"`
	RabbitMQ rabbitMQConfig `mapstructure:"rabbitmq" toml:"rabbitmq" yaml:"rabbitmq" json:"rabbitmq"`
}

type azureConfig struct {
	BusConnectionString string `mapstructure:"bus_connection_string" toml:"bus_connection_string" yaml:"bus_connection_string" json:"bus_connection_string"`
	Topic               string `mapstructure:"topic" toml:"topic" yaml:"topic" json:"topic"`
}

type rabbitMQConfig struct {
	Host     string `mapstructure:"host" toml:"host" yaml:"host" json:"host"`
	Port     uint   `mapstructure:"port" toml:"port" yaml:"port" json:"port"`
	User     string `mapstructure:"user" toml:"user" yaml:"user" json:"user"`
	Password string `mapstructure:"password" toml:"password" yaml:"password" json:"password"`
	Vhost    string `mapstructure:"vhost" toml:"vhost" yaml:"vhost" json:"vhost"`
	Exchange string `mapstructure:"exchange" toml:"exchange" yaml:"exchange" json:"exchange"`
}

type sourceConfigs struct {
	Pixiv SourceConfig `mapstructure:"pixiv" toml:"pixiv" yaml:"pixiv" json:"pixiv"`
}

type SourceConfig struct {
	Enable   bool     `mapstructure:"enable" toml:"enable" yaml:"enable" json:"enable"`
	URLs     []string `mapstructure:"urls" toml:"urls" yaml:"urls" json:"urls"`
	Interval uint     `mapstructure:"interval" toml:"interval" yaml:"interval" json:"interval"`
}

var Cfg *tomlConfig

func init() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AddConfigPath("..")
	viper.SetConfigType("toml")
	viper.SetEnvPrefix("MANYACG")
	viper.AutomaticEnv()
	viper.SetDefault("app.debug", false)
	viper.SetDefault("log.level", "info")
	viper.SetDefault("log.file_path", "./logs/collector.log")
	viper.SetDefault("log.backup_num", 7)
	viper.SetDefault("sources.pixiv.enable", false)
	viper.SetDefault("sender.type", "rabbitmq")
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
