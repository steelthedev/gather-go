package db

import (
	"gather-go/package/accounts/models"
	"gather-go/package/room/room_models"

	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(url string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {

		log.Fatalln(err)

	}

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&room_models.Room{})

	return db

}
