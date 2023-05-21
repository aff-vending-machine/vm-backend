package rabbitmq

import (
	"fmt"
	"vm-backend/configs"
	"vm-backend/pkg/boot"

	"github.com/rs/zerolog/log"
)

type Service struct {
	*Connection
	configs.RabbitMQConfig
}

func New(cfg configs.RabbitMQConfig) *Service {
	url := fmt.Sprintf(
		"%s://%s:%s@%s:%s/%s",
		cfg.Protocol,
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Path,
	)
	conn, err := Dial(url)
	if err != nil {
		log.Fatal().Err(err).Str("url", url).Interface("config", cfg).Msg("unable to connect to rabbitmq")
		boot.Signal.Stop()
	}

	return &Service{
		conn,
		cfg,
	}
}
