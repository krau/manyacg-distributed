package config

type appConfig struct {
	Debug bool        `mapstructure:"debug" toml:"debug" yaml:"debug" json:"debug"`
	Redis redisConfig `mapstructure:"redis" toml:"redis" yaml:"redis" json:"redis"`
}

type redisConfig struct {
	Addr     string `mapstructure:"addr" toml:"addr" yaml:"addr" json:"addr"`
	Password string `mapstructure:"password" toml:"password" yaml:"password" json:"password"`
	DB       int    `mapstructure:"db" toml:"db" yaml:"db" json:"db"`
}
