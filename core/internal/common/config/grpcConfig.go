package config

type grpcConfig struct {
	Address  string `mapstructure:"address" toml:"address" yaml:"address" json:"address"`
	CertFile string `mapstructure:"cert" toml:"cert" yaml:"cert" json:"cert"`
	KeyFile  string `mapstructure:"key" toml:"key" yaml:"key" json:"key"`
	CaFile   string `mapstructure:"ca" toml:"ca" yaml:"ca" json:"ca"`
}
