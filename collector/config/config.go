package config

import (
	"github.com/spf13/viper"
)

type tomlConfig struct {
	App     appConfig    `mapstructure:"app" toml:"app" yaml:"app" json:"app"`
	Sources sourceConfig `mapstructure:"sources" toml:"sources" yaml:"sources" json:"sources"`
	Core    core
}

type appConfig struct {
	Log   logConfig `mapstructure:"log" toml:"log" yaml:"log" json:"log"`
	Debug bool      `mapstructure:"debug" toml:"debug" yaml:"debug" json:"debug"`
}

type logConfig struct {
	Level     string `mapstructure:"level" toml:"level" yaml:"level" json:"level"`
	FilePath  string `mapstructure:"file_path" toml:"file_path" yaml:"file_path" json:"file_path"`
	BackupNum uint   `mapstructure:"backup_num" toml:"backup_num" yaml:"backup_num" json:"backup_num"`
}

type sourceConfig struct {
	Pixiv source `mapstructure:"pixiv" toml:"pixiv" yaml:"pixiv" json:"pixiv"`
}

type source struct {
	Enable   bool   `mapstructure:"enable" toml:"enable" yaml:"enable" json:"enable"`
	URL      string `mapstructure:"url" toml:"url" yaml:"url" json:"url"`
	Interval uint   `mapstructure:"interval" toml:"interval" yaml:"interval" json:"interval"`
}

type core struct {
	URL      string `mapstructure:"url" toml:"url" yaml:"url" json:"url"`
	Username string `mapstructure:"username" toml:"username" yaml:"username" json:"username"`
	Password string `mapstructure:"password" toml:"password" yaml:"password" json:"password"`
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
	viper.SetDefault("app.log.file_path", "./logs/collector.log")
	viper.SetDefault("app.log.backup_num", 7)
	viper.SetDefault("sources.pixiv.enable", false)
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
