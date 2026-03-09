package db

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

var _ DB = (*PostrgresDb)(nil)

type PostrgresDb struct {
	postrgres *pgxpool.Pool
}

func NewPostgresDb(postgresURI string, postgresConnections int) (*PostrgresDb, error) {
	poolConfig, err := pgxpool.ParseConfig(postgresURI)
	if err != nil {
		return nil, err
	}

	poolConfig.MaxConns = int32(postgresConnections)
	postgres, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		return nil, err
	}

	return &PostrgresDb{
		postrgres: postgres,
	}, nil
}

func (p *PostrgresDb) QueryRow(ctx context.Context, sql string, args ...any) pgx.Row {
	return p.postrgres.QueryRow(ctx, sql, args...)
}

func (p *PostrgresDb) Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error) {
	return p.postrgres.Query(ctx, sql, args...)
}

func (p *PostrgresDb) Exec(ctx context.Context, sql string, args ...any) (pgconn.CommandTag, error) {
	return p.postrgres.Exec(ctx, sql, args...)
}

func (p *PostrgresDb) Begin(ctx context.Context) (pgx.Tx, error) {
	return p.postrgres.Begin(ctx)
}

func (p *PostrgresDb) Close() {
	p.postrgres.Close()
}
