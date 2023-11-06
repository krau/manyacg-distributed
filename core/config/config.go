package config

import (
	"github.com/spf13/viper"
)

type tomlConfig struct {
	App       appConfig       `mapstructure:"app" toml:"app" yaml:"app" json:"app"`
	Database  databaseConfig  `mapstructure:"database" toml:"database" yaml:"database" json:"database"`
	Log       logConfig       `mapstructure:"log" toml:"log" yaml:"log" json:"log"`
	Messenger messengerConfig `mapstructure:"messenger" toml:"messenger" yaml:"messenger" json:"messenger"`
	API       apiConfig       `mapstructure:"api" toml:"api" yaml:"api" json:"api"`
}

type appConfig struct {
	Debug       bool   `mapstructure:"debug" toml:"debug" yaml:"debug" json:"debug"`
	ExtProcess  bool   `mapstructure:"ext_process" toml:"ext_process" yaml:"ext_process" json:"ext_process"`
	ImagePrefix string `mapstructure:"image_prefix" toml:"image_prefix" yaml:"image_prefix" json:"image_prefix"`
	Address     string `mapstructure:"address" toml:"address" yaml:"address" json:"address"`
	CertFile    string `mapstructure:"cert" toml:"cert" yaml:"cert" json:"cert"`
	KeyFile     string `mapstructure:"key" toml:"key" yaml:"key" json:"key"`
	CaFile      string `mapstructure:"ca" toml:"ca" yaml:"ca" json:"ca"`
}

type logConfig struct {
	Level     string `mapstructure:"level" toml:"level" yaml:"level" json:"level"`
	FilePath  string `mapstructure:"file_path" toml:"file_path" yaml:"file_path" json:"file_path"`
	BackupNum uint   `mapstructure:"backup_num" toml:"backup_num" yaml:"backup_num" json:"backup_num"`
}

type messengerConfig struct {
	Type     string         `mapstructure:"type" toml:"type" yaml:"type" json:"type"`
	Azure    azureConfig    `mapstructure:"azure" toml:"azure" yaml:"azure" json:"azure"`
	RabbitMQ rabbitMQConfig `mapstructure:"rabbitmq" toml:"rabbitmq" yaml:"rabbitmq" json:"rabbitmq"`
}

type azureConfig struct {
	BusConnectionString string `mapstructure:"bus_connection_string" toml:"bus_connection_string" yaml:"bus_connection_string" json:"bus_connection_string"`
	SubTopic            string `mapstructure:"sub_topic" toml:"sub_topic" yaml:"sub_topic" json:"sub_topic"`
	Subscription        string `mapstructure:"subscription" toml:"subscription" yaml:"subscription" json:"subscription"`
	PubTopic            string `mapstructure:"pub_topic" toml:"pub_topic" yaml:"pub_topic" json:"pub_topic"`
}

type rabbitMQConfig struct {
	Host        string `mapstructure:"host" toml:"host" yaml:"host" json:"host"`
	Port        int    `mapstructure:"port" toml:"port" yaml:"port" json:"port"`
	User        string `mapstructure:"user" toml:"user" yaml:"user" json:"user"`
	Password    string `mapstructure:"password" toml:"password" yaml:"password" json:"password"`
	Vhost       string `mapstructure:"vhost" toml:"vhost" yaml:"vhost" json:"vhost"`
	SubExchange string `mapstructure:"sub_exchange" toml:"sub_exchange" yaml:"sub_exchange" json:"sub_exchange"`
	SubQueue    string `mapstructure:"sub_queue" toml:"sub_queue" yaml:"sub_queue" json:"sub_queue"`
	PubExchange string `mapstructure:"pub_exchange" toml:"pub_exchange" yaml:"pub_exchange" json:"pub_exchange"`
}

type databaseConfig struct {
	Host     string `mapstructure:"host" toml:"host" yaml:"host" json:"host"`
	Port     int    `mapstructure:"port" toml:"port" yaml:"port" json:"port"`
	User     string `mapstructure:"user" toml:"user" yaml:"user" json:"user"`
	Password string `mapstructure:"password" toml:"password" yaml:"password" json:"password"`
	Database string `mapstructure:"database" toml:"database" yaml:"database" json:"database"`
}

type apiConfig struct {
	Enable bool   `mapstructure:"enable" toml:"enable" yaml:"enable" json:"enable"`
	Address string `mapstructure:"address" toml:"address" yaml:"address" json:"address"`
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
	viper.SetDefault("app.ext_process", false)
	viper.SetDefault("app.image_prefix", "./")
	viper.SetDefault("app.address", "0.0.0.0:39010")
	viper.SetDefault("log.level", "info")
	viper.SetDefault("log.file_path", "./logs/storage.log")
	viper.SetDefault("log.backup_num", 7)
	viper.SetDefault("messenger.type", "rabbitmq")
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
