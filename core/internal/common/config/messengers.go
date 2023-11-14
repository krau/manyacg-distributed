package config

type messengerConfig struct {
	Type     string         `mapstructure:"type" toml:"type" yaml:"type" json:"type"`
	Azure    azureConfig    `mapstructure:"azure" toml:"azure" yaml:"azure" json:"azure"`
	RabbitMQ rabbitMQConfig `mapstructure:"rabbitmq" toml:"rabbitmq" yaml:"rabbitmq" json:"rabbitmq"`
}

type azureConfig struct {
	BusConnectionString string `mapstructure:"bus_connection_string" toml:"bus_connection_string" yaml:"bus_connection_string" json:"bus_connection_string"`
	SubTopic            string `mapstructure:"sub_topic" toml:"sub_topic" yaml:"sub_topic" json:"sub_topic"`
	Subscription        string `mapstructure:"subscription" toml:"subscription" yaml:"subscription" json:"subscription"`
	PubTopic            string `mapstructure:"pub_topic" toml:"pub_topic" yaml:"pub_topic" json:"pub_topic"`
}

type rabbitMQConfig struct {
	Host        string `mapstructure:"host" toml:"host" yaml:"host" json:"host"`
	Port        int    `mapstructure:"port" toml:"port" yaml:"port" json:"port"`
	User        string `mapstructure:"user" toml:"user" yaml:"user" json:"user"`
	Password    string `mapstructure:"password" toml:"password" yaml:"password" json:"password"`
	Vhost       string `mapstructure:"vhost" toml:"vhost" yaml:"vhost" json:"vhost"`
	SubExchange string `mapstructure:"sub_exchange" toml:"sub_exchange" yaml:"sub_exchange" json:"sub_exchange"`
	SubQueue    string `mapstructure:"sub_queue" toml:"sub_queue" yaml:"sub_queue" json:"sub_queue"`
	PubExchange string `mapstructure:"pub_exchange" toml:"pub_exchange" yaml:"pub_exchange" json:"pub_exchange"`
}
