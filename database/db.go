package database

import (
	"bivrost_task2/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "10.5.0.5"
	user     = "postgres"
	password = "postgres"
	dbPort   = "5432"
	dbName   = "orderDB"
	db       *gorm.DB
	err      error
)

func StartDB() {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbName, dbPort)
	dsn := config
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Database connection failed: ", err)
	}

	fmt.Println("Database connection success")
	db.Debug().AutoMigrate(models.Item{}, models.Orders{})

}

func GetDB() *gorm.DB {
	return db
}
