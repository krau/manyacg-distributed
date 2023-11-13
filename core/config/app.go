package config


type appConfig struct {
	Debug       bool   `mapstructure:"debug" toml:"debug" yaml:"debug" json:"debug"`
}