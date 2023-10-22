package config

import "github.com/spf13/viper"

type tomlConfig struct {
	App        appConfig        `mapstructure:"app" toml:"app" yaml:"app" json:"app"`
	Storages   storageConfigs   `mapstructure:"storages" toml:"storages" yaml:"storages" json:"storages"`
	Log        logConfig        `mapstructure:"log" toml:"log" yaml:"log" json:"log"`
	Subscriber subscriberConfig `mapstructure:"subscriber" toml:"subscriber" yaml:"subscriber" json:"subscriber"`
}

type appConfig struct {
	Debug      bool   `mapstructure:"debug" toml:"debug" yaml:"debug" json:"debug"`
	Sleep      uint   `mapstructure:"sleep" toml:"sleep" yaml:"sleep" json:"sleep"`
	GrpcAddr   string `mapstructure:"grpc_addr" toml:"grpc_addr" yaml:"grpc_addr" json:"grpc_addr"`
	ServerName string `mapstructure:"server_name" toml:"server_name" yaml:"server_name" json:"server_name"`
	CertFile   string `mapstructure:"cert" toml:"cert" yaml:"cert" json:"cert"`
	KeyFile    string `mapstructure:"key" toml:"key" yaml:"key" json:"key"`
	CaFile     string `mapstructure:"ca" toml:"ca" yaml:"ca" json:"ca"`
}

type logConfig struct {
	Level     string `mapstructure:"level" toml:"level" yaml:"level" json:"level"`
	FilePath  string `mapstructure:"file_path" toml:"file_path" yaml:"file_path" json:"file_path"`
	BackupNum uint   `mapstructure:"backup_num" toml:"backup_num" yaml:"backup_num" json:"backup_num"`
}

type storageConfigs struct {
	Local    storageLocalConfig    `mapstructure:"local" toml:"local" yaml:"local" json:"local"`
	Telegram storageTelegramConfig `mapstructure:"telegram" toml:"telegram" yaml:"telegram" json:"telegram"`
}

type storageLocalConfig struct {
	Enable bool   `mapstructure:"enable" toml:"enable" yaml:"enable" json:"enable"`
	Dir    string `mapstructure:"dir" toml:"dir" yaml:"dir" json:"dir"`
}

type storageTelegramConfig struct {
	Enable   bool   `mapstructure:"enable" toml:"enable" yaml:"enable" json:"enable"`
	Token    string `mapstructure:"token" toml:"token" yaml:"token" json:"token"`
	ChatId   int64  `mapstructure:"chat_id" toml:"chat_id" yaml:"chat_id" json:"chat_id"`
	Username string `mapstructure:"username" toml:"username" yaml:"username" json:"username"`
}

type subscriberConfig struct {
	Azure azureConfig `mapstructure:"azure" toml:"azure" yaml:"azure" json:"azure"`
}

type azureConfig struct {
	BusConnectionString string `mapstructure:"bus_connection_string" toml:"bus_connection_string" yaml:"bus_connection_string" json:"bus_connection_string"`
	SubTopic            string `mapstructure:"sub_topic" toml:"sub_topic" yaml:"sub_topic" json:"sub_topic"`
	Subscription        string `mapstructure:"subscription" toml:"subscription" yaml:"subscription" json:"subscription"`
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
	viper.SetDefault("app.sleep", 5)
	viper.SetDefault("log.level", "info")
	viper.SetDefault("log.file_path", "./logs/storage.log")
	viper.SetDefault("log.backup_num", 7)
	viper.SetDefault("storages.local.enable", false)
	viper.SetDefault("storages.local.dir", "./pictures")
	viper.SetDefault("storages.telegram.enable", false)
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
