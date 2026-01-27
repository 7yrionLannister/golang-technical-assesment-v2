package db

import (
	"database/sql"
	"os"

	"github.com/7yrionLannister/golang-technical-assesment-v2/pkg/env"
	"github.com/7yrionLannister/golang-technical-assesment-v2/pkg/log"
	"github.com/7yrionLannister/golang-technical-assesment-v2/pkg/util"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/v5/stdlib"
)

// Run the golang-migrate migrations defined in the db.migrations folder
func MigrateUp() error {
	log.L.Debug("Applying migrations...")
	sqlDB, err := sql.Open("pgx", env.Env.DataBaseUrl)
	if err != nil {
		log.L.Error("Failed to open database", "error", err)
		os.Exit(1)
	}
	defer sqlDB.Close()

	// Run migrations with golang-migrate
	m, err := migrate.New(
		"file://./../../pkg/config/db.migrations",
		env.Env.DataBaseUrl,
	)
	if err != nil {
		return util.HandleError(err, "Failed to create migrate instance")
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return util.HandleError(err, "Failed to apply migrations")
	}

	log.L.Debug("Migrations applied successfully!")
	return nil
}
