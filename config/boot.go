package config

type BootConfig struct {
	App        AppConfig        `mapstructure:"APP"`
	Fiber      FiberConfig      `mapstructure:"FIBER"`
	PostgreSQL PostgreSQLConfig `mapstructure:"POSTGRESQL"`
	RabbitMQ   RabbitMQConfig   `mapstructure:"RABBITMQ"`
}
