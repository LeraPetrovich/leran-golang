package storage

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"example.com/learn/pkg/publishPlatformMusic/core"
	"go.uber.org/zap"
)

const fastQueryTimeout = 3 * time.Second

type storage struct {
	logger   *zap.Logger
	postgres *pgxpool.Pool
}

type StorageConf struct {
	PostgresConnections int
}

var _ core.AppStorage = (*storage)(nil)

func New(logger *zap.Logger, postgresURI string, conf StorageConf) (*storage, error) {
	poolConfig, err := pgxpool.ParseConfig(postgresURI)

	if err != nil {
		return nil, err
	}

	poolConfig.MaxConns = int32(conf.PostgresConnections)

	postgres, err := pgxpool.NewWithConfig(context.TODO(), poolConfig)
	if err != nil {
		return nil, err
	}

	return &storage{
		logger:   logger,
		postgres: postgres,
	}, nil
}
