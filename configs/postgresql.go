package configs

type PostgreSQLConfig struct {
	LogLevel int    `default:"1" mapstructure:"LOG_LEVEL"`
	Pool     bool   `default:"false" mapstructure:"POOL"`
	Username string `default:"" mapstructure:"USERNAME"`
	Password string `default:"" mapstructure:"PASSWORD"`
	Host     string `default:"localhost" mapstructure:"HOST"`
	Port     int    `default:"5432" mapstructure:"PORT"`
	Database string `default:"postgresdb" mapstructure:"DATABASE"`
	SSLMode  string `default:"disable" mapstructure:"SSLMODE"`
	TimeZone string `default:"Asia/Bangkok" mapstructure:"TIMEZONE"`
}
