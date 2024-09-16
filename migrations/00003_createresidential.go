package migrations

import (
	"app/database"
	"app/models"
	"context"
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upCreateresidential, downCreateresidential)
}

func upCreateresidential(ctx context.Context, tx *sql.Tx) error {
	return database.DB_MIGRATOR.CreateTable(&models.ResidentialDetails{})
}

func downCreateresidential(ctx context.Context, tx *sql.Tx) error {
	return database.DB_MIGRATOR.DropTable(&models.ResidentialDetails{})
}
