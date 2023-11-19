package config

type redisConfig struct {
	URL      string `mapstructure:"url" toml:"url" yaml:"url" json:"url"`
	CacheTTL int    `mapstructure:"cache_ttl" toml:"cache_ttl" yaml:"cache_ttl" json:"cache_ttl"`
}
