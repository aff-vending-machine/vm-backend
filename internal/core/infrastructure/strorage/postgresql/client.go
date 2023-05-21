package postgresql

import (
	"fmt"
	"time"

	"vm-backend/configs"
	"vm-backend/pkg/boot"

	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Client struct {
	*gorm.DB
	configs.PostgreSQLConfig
}

func New(cfg configs.PostgreSQLConfig) *Client {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		cfg.Host,
		cfg.Username,
		cfg.Password,
		cfg.Database,
		cfg.Port,
		cfg.SSLMode,
		cfg.TimeZone,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		PrepareStmt:            cfg.Pool,
		SkipDefaultTransaction: true,
		Logger:                 logger.Default.LogMode(logger.LogLevel(cfg.LogLevel)),
	})
	if err != nil {
		log.Error().Err(err).Interface("config", cfg).Msg("unable to connect to database")
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Error().Err(err).Interface("config", cfg).Msg("unable to connect to database")
		boot.Signal.Stop()
	}

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(5)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(50)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(24 * time.Hour)

	err = sqlDB.Ping()
	if err != nil {
		log.Error().Err(err).Str("dsn", dsn).Interface("config", cfg).Msg("unable to connect to database")
		boot.Signal.Stop()
	}

	return &Client{
		db,
		cfg,
	}
}
