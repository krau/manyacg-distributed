package config

type appConfig struct {
	Debug    bool   `mapstructure:"debug" toml:"debug" yaml:"debug" json:"debug"`
	CacheDir string `mapstructure:"cache_dir" toml:"cache_dir" yaml:"cache_dir" json:"cache_dir"`
}
