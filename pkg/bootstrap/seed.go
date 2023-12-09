package bootstrap

import (
	"ginblog/internal/database/seeder"
	"ginblog/pkg/config"
	"ginblog/pkg/database"
)

func Seed() {
	config.Set()
	database.Connect()
	seeder.Seed()
}
