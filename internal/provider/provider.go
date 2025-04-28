package provider

import (
	"learnyscape-backend-mono/internal/config"
	"learnyscape-backend-mono/internal/data"
	"learnyscape-backend-mono/pkg/database"

	"github.com/jmoiron/sqlx"
)

var (
	db        *sqlx.DB
	dataStore data.DataStore
)

func BootstrapGlobal(cfg *config.Config) {
	db = database.NewPostgres((*database.PostgresOptions)(cfg.Postgres))
	dataStore = data.NewDataStore(db)
}
