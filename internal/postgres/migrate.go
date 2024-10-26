package postgres

import (
	"context"
	"database/sql"
	"embed"
	"github.com/pressly/goose/v3"

	"github.com/sashaaro/gophkeeper/internal/log"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

func RunMigrations(ctx context.Context, db *sql.DB) error {
	goose.SetBaseFS(embedMigrations)

	log.Info("DB Migration: start")
	if err := goose.SetDialect("postgres"); err != nil {
		return err
	}

	if err := goose.UpContext(ctx, db, "migrations"); err != nil {
		return err
	}
	log.Info("DB Migration: success")

	return nil
}
