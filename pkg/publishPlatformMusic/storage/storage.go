package storage

import (
	"example.com/learn/pkg/publishPlatformMusic/core"
	"example.com/learn/pkg/publishPlatformMusic/db"
	"go.uber.org/zap"
	"time"
)

const fastQueryTimeout = 3 * time.Second

type storage struct {
	logger *zap.Logger
	db     db.DB
}

var _ core.AppStorage = (*storage)(nil)

func New(logger *zap.Logger, database db.DB) *storage {
	return &storage{
		logger: logger,
		db:     database,
	}
}
