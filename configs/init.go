package configs

import (
	"os"
	"reflect"
	"strings"

	"github.com/creasty/defaults"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

type Config struct {
	App        AppConfig        `mapstructure:"APP"`
	Fiber      FiberConfig      `mapstructure:"FIBER"`
	PostgreSQL PostgreSQLConfig `mapstructure:"POSTGRESQL"`
	RabbitMQ   RabbitMQConfig   `mapstructure:"RABBITMQ"`
	WebSocket  WebSocketConfig  `mapstructure:"WEBSOCKET"`
}

// Init creates a new config service.
func Init(fallback string) Config {
	filename := fallback
	if value, ok := os.LookupEnv("APP_ENV"); ok {
		filename = value
	}

	var out Config
	if err := defaults.Set(&out); err != nil {
		log.Panic().Err(err).Msgf("set default config failed")
	}

	v := viper.New()
	v.SetConfigName(filename)
	v.AddConfigPath(".")
	v.SetConfigType("yaml")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))

	// read config file, ignore error when not found.
	if err := v.ReadInConfig(); err != nil {
		log.Warn().Err(err).Msgf("read config file failed")
	}

	// set environment variables as overrides.
	bindEnvs(v, Config{})
	v.AutomaticEnv()

	// set config form file or env
	if err := v.Unmarshal(&out); err != nil {
		log.Warn().Err(err).Msgf("unmarshal config failed")
	}

	return out
}

func bindEnvs(vp *viper.Viper, iface interface{}, parts ...string) {
	ifv := reflect.ValueOf(iface)
	ift := reflect.TypeOf(iface)
	for i := 0; i < ift.NumField(); i++ {
		v := ifv.Field(i)
		t := ift.Field(i)
		tv, ok := t.Tag.Lookup("mapstructure")
		if !ok {
			continue
		}
		switch v.Kind() {
		case reflect.Struct:
			bindEnvs(vp, v.Interface(), append(parts, tv)...)
		default:
			vp.BindEnv(strings.Join(append(parts, tv), "."))
		}
	}
}

func (cfg *Config) Print() {
	log.Debug().Str("version", cfg.App.Version).Msg("configuration")
	log.Debug().Interface("App", cfg.App).Send()
	log.Debug().Interface("Fiber", cfg.Fiber).Send()
	log.Debug().Interface("PostgreSQL", cfg.PostgreSQL).Send()
	log.Debug().Interface("RabbitMQ", cfg.RabbitMQ).Send()
}
