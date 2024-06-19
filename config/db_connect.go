package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func NewDbConnection() *gorm.DB {
	db, err := gorm.Open(postgres.Open(EnvPostgresURI()))
	if err != nil {
		log.Println("#ERROR " + err.Error())
	}
	return db
}
