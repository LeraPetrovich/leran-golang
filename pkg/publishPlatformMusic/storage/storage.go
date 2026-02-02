package storage

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type storage struct {
	logger   *zap.Logger
	postgres *pgxpool.Pool
}

type StorageConf struct {
	PostgresConnections int
}

func New(logger *zap.Logger, postgresURI string, conf StorageConf) (*storage, error) {
	poolConfig, err := pgxpool.ParseConfig(postgresURI)

	if err != nil {
		return nil, err
	}

	poolConfig.MaxConns = int32(conf.PostgresConnections)

	postgres, err := pgxpool.NewWithConfig(context.TODO(), poolConfig)

	return &storage{
		logger:   logger,
		postgres: postgres,
	}, nil
}
