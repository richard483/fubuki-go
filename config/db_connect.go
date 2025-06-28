package config

import (
	"log/slog"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDbConnection() *gorm.DB {
	db, err := gorm.Open(postgres.Open(EnvPostgresURI()))
	if err != nil {
		slog.Error("#NewDbConnection - error on opening Gorm connection", "error", err.Error())
		panic(err)
	}
	return db
}
