//go:build migrate
// +build migrate

package main

import (
	"fubuki-go/config"
	"fubuki-go/model"
	"log"
)

func init() {
	db := config.NewDbConnection()
	log.Println("Masuk migrate History")
	err := db.AutoMigrate(&model.History{})

	if err != nil {
		log.Fatalln(err)
	}
}
