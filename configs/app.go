package configs

type AppConfig struct {
	ENV       string       `default:"production" mapstructure:"ENV"`
	LogLevel  int          `default:"0" mapstructure:"LOG_LEVEL"`
	Version   string       `default:"0.0.0" mapstructure:"VERSION"`
	Migration bool         `default:"false" mapstructure:"MIGRATION"`
	BCrypt    BCryptConfig `mapstructure:"BCRYPT"`
	JWT       JWTConfig    `mapstructure:"JWT"`
}

type BCryptConfig struct {
	Salt int `default:"6" mapstructure:"SALT"`
}

type JWTConfig struct {
	Reference1      string `default:"reference1" mapstructure:"REFERENCE1"`
	Reference2      string `default:"reference2" mapstructure:"REFERENCE2"`
	Audience        string `default:"vm-backend" mapstructure:"AUDIENCE"`
	Issuer          string `default:"vm" mapstructure:"ISSUER"`
	AuthorizedParty string `default:"none" mapstructure:"AUTHORIZED_PARTY"`
}
