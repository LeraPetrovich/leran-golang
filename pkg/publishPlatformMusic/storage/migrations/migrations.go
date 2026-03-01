package migrations

import (
	"database/sql"
	"embed"

	"github.com/pressly/goose/v3"
	_ "github.com/jackc/pgx/v5/stdlib"
)

//go:embed *.sql
var embedMigrations embed.FS

func Run(postgresURI string) error {
	db, err := sql.Open("pgx", postgresURI)
	if err != nil {
		return err
	}
	defer db.Close()

	goose.SetBaseFS(embedMigrations)

	if err := goose.SetDialect("postgres"); err != nil {
		return err
	}

	if err := goose.Up(db, "."); err != nil {
		return err
	}
	return nil
}
