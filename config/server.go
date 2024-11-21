package config

type Server struct {
	Port     int    `mapstructure:"port"`
	DBType   string `mapstructure:"db-type"`
	UseRedis bool   `mapstructure:"use-redis"`
}
