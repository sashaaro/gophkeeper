package postgres

import (
	"context"
	"embed"
	"github.com/jmoiron/sqlx"
	"github.com/pressly/goose/v3"

	"github.com/sashaaro/gophkeeper/internal/log"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

func RunMigrations(ctx context.Context, db *sqlx.DB) error {
	goose.SetBaseFS(embedMigrations)

	log.Info("DB Migration: start")
	if err := goose.SetDialect("postgres"); err != nil {
		return err
	}

	if err := goose.UpContext(ctx, db.DB, "migrations"); err != nil {
		return err
	}
	log.Info("DB Migration: success")

	return nil
}
