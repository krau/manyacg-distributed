package config

type middlewareConfig struct {
	MQ    mqConfig    `mapstructure:"messenger" toml:"messenger" yaml:"messenger" json:"messenger"`
	Redis redisConfig `mapstructure:"redis" toml:"redis" yaml:"redis" json:"redis"`
}
