package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Conn() {
	URI := "user=root password=root dbname=root port=5432 host=localhost sslmode=disable"
	DB, err = gorm.Open(postgres.Open(URI))
	if err != nil {
		log.Panic("Error:" + err.Error())
	}
}
