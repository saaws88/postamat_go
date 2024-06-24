package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"saaws88/model"
)

var DB *gorm.DB

func ConnectToDb() {

	dsn := "host=localhost port=port user=postgres dbname=postamat_go password=password sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Невозможно подключиться к БД")
	}

	db.AutoMigrate(&model.Postamat{})

	DB = db

}
