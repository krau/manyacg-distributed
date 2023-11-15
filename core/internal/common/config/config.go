package config

import (
	"github.com/spf13/viper"
)

type tomlConfig struct {
	App        appConfig        `mapstructure:"app" toml:"app" yaml:"app" json:"app"`
	Database   databaseConfig   `mapstructure:"database" toml:"database" yaml:"database" json:"database"`
	Log        logConfig        `mapstructure:"log" toml:"log" yaml:"log" json:"log"`
	Middleware middlewareConfig `mapstructure:"middleware" toml:"middleware" yaml:"middleware" json:"middleware"`
	Processor  processorConfig  `mapstructure:"processor" toml:"processor" yaml:"processor" json:"processor"`
	GRPC       grpcConfig       `mapstructure:"grpc" toml:"grpc" yaml:"grpc" json:"grpc"`
	API        apiConfig        `mapstructure:"api" toml:"api" yaml:"api" json:"api"`
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
