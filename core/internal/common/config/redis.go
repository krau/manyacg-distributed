package config

type redisConfig struct {
	Addr     string `mapstructure:"addr" toml:"addr" yaml:"addr" json:"addr"`
	Password string `mapstructure:"password" toml:"password" yaml:"password" json:"password"`
	DB       int    `mapstructure:"db" toml:"db" yaml:"db" json:"db"`
	CacheTTL int    `mapstructure:"cache_ttl" toml:"cache_ttl" yaml:"cache_ttl" json:"cache_ttl"`
}
