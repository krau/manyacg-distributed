package config

type apiConfig struct {
	Enable bool   `mapstructure:"enable" toml:"enable" yaml:"enable" json:"enable"`
	Address string `mapstructure:"address" toml:"address" yaml:"address" json:"address"`
}