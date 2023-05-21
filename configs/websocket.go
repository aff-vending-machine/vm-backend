package configs

type WebSocketConfig struct {
	Host string `mapstructure:"HOST" default:"localhost"`
	Port string `mapstructure:"PORT" default:"8080"`
}
