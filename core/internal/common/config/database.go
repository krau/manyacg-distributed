package config

type databaseConfig struct {
	Host     string `mapstructure:"host" toml:"host" yaml:"host" json:"host"`
	Port     int    `mapstructure:"port" toml:"port" yaml:"port" json:"port"`
	User     string `mapstructure:"user" toml:"user" yaml:"user" json:"user"`
	Password string `mapstructure:"password" toml:"password" yaml:"password" json:"password"`
	Database string `mapstructure:"database" toml:"database" yaml:"database" json:"database"`
}
