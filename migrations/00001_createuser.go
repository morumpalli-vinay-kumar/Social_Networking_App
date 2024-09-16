package migrations

import (
	"app/database"
	"app/models"
	"context"
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upCreateuser, downCreateuser)
}

func upCreateuser(ctx context.Context, tx *sql.Tx) error {
	return database.DB_MIGRATOR.CreateTable(&models.UserDetails{})
}

func downCreateuser(ctx context.Context, tx *sql.Tx) error {
	return database.DB_MIGRATOR.DropTable(&models.UserDetails{})
}
