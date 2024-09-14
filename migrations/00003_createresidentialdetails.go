package migrations

import (
	"app/database"
	"app/models"
	"context"
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upCreateresidentialdetails, downCreateresidentialdetails)
}

func upCreateresidentialdetails(ctx context.Context, tx *sql.Tx) error {
	return database.DB_MIGRATOR.CreateTable(&models.ResidentialDetails{})
}

func downCreateresidentialdetails(ctx context.Context, tx *sql.Tx) error {
	return database.DB_MIGRATOR.DropTable(&models.ResidentialDetails{})
}
