package migrations

import (
	"app/database"
	"app/models"
	"context"
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upCreateaddress, downCreateaddress)
}

func upCreateaddress(ctx context.Context, tx *sql.Tx) error {
	return database.DB_MIGRATOR.CreateTable(&models.Address{})
}

func downCreateaddress(ctx context.Context, tx *sql.Tx) error {
	return database.DB_MIGRATOR.DropTable(&models.Address{})
}
