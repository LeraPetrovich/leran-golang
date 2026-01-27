package storage

import (
	"context"
	"time"

	learnCore "example.com/learn/pkg/learn/core"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

const fastQueryTimeout = 3 * time.Second

type storage struct {
	logger   *zap.Logger
	postgres *pgxpool.Pool
}

type Config struct {
	PostgresConnections int
}

var _ learnCore.AppStorage = (*storage)(nil)

func New(logger *zap.Logger, postgresURI string, config Config) (*storage, error) {
	poolConfig, err := pgxpool.ParseConfig(postgresURI)

	if err != nil {
		return nil, err
	}

	poolConfig.MaxConns = int32(config.PostgresConnections)

	postgres, err := pgxpool.NewWithConfig(context.TODO(), poolConfig)

	if err != nil {
		return nil, err
	}
	s := storage{
		logger:   logger,
		postgres: postgres,
	}
	return &s, nil
}
