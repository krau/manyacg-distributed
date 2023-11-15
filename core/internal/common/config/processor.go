package config

type processorConfig struct {
	EnableExt     bool           `mapstructure:"enable_ext" toml:"enable_ext" yaml:"enable_ext" json:"enable_ext"`
	MaxConcurrent int            `mapstructure:"max_concurrent" toml:"max_concurrent" yaml:"max_concurrent" json:"max_concurrent"`
	Download      downloadConfig `mapstructure:"download" toml:"download" yaml:"download" json:"download"`
	Save          saveConfig     `mapstructure:"save" toml:"save" yaml:"save" json:"save"`
}

type downloadConfig struct {
}

type saveConfig struct {
	Type   string       `mapstructure:"type" toml:"type" yaml:"type" json:"type"`
	Local  localConfig  `mapstructure:"local" toml:"local" yaml:"local" json:"local"`
	Webdav webdavConfig `mapstructure:"webdav" toml:"webdav" yaml:"webdav" json:"webdav"`
}

type localConfig struct {
	Path string `mapstructure:"dir" toml:"dir" yaml:"dir" json:"dir"`
}

type webdavConfig struct {
	URL      string `mapstructure:"url" toml:"url" yaml:"url" json:"url"`
	Username string `mapstructure:"username" toml:"username" yaml:"username" json:"username"`
	Password string `mapstructure:"password" toml:"password" yaml:"password" json:"password"`
	Path     string `mapstructure:"path" toml:"path" yaml:"path" json:"path"`
	CacheTTL int    `mapstructure:"cache_ttl" toml:"cache_ttl" yaml:"cache_ttl" json:"cache_ttl"`
}
