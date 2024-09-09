package database

import (
	"database/sql"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var GORM_DB *gorm.DB
var SQL_DB *sql.DB
var DB_MIGRATOR gorm.Migrator

func ConnectToDatabase(dbURL string) error {
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err == nil {
		GORM_DB = db
		SQL_DB, _ = db.DB()
		DB_MIGRATOR = db.Migrator()
	}
	return err
}
