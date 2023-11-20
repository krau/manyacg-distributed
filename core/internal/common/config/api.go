package config

type apiConfig struct {
	Enable           bool   `mapstructure:"enable" toml:"enable" yaml:"enable" json:"enable"`
	Address          string `mapstructure:"address" toml:"address" yaml:"address" json:"address"`
	EnableRedisCache bool   `mapstructure:"enable_redis_cache" toml:"enable_redis_cache" yaml:"enable_redis_cache" json:"enable_redis_cache"`
}
