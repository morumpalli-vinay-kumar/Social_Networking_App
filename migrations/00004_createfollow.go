package migrations

import (
	"app/models"
	"context"
	"database/sql"

	"github.com/pressly/goose/v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func init() {
	goose.AddMigrationContext(upCreatefollow, downCreatefollow)
}

func upCreatefollow(ctx context.Context, tx *sql.Tx) error {
	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: tx,
	}), &gorm.Config{})
	if err != nil {
		return err
	}
	return db.Migrator().CreateTable(&models.Follow{})
}

func downCreatefollow(ctx context.Context, tx *sql.Tx) error {
	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: tx,
	}), &gorm.Config{})
	if err != nil {
		return err
	}
	return db.Migrator().CreateTable(&models.Follow{})
}
