package bootstrap

import (
	"ginblog/internal/database/migration"
	"ginblog/pkg/config"
	"ginblog/pkg/database"
)

func Migrate() {
	config.Set()
	database.Connect()
	migration.Migrate()
}
