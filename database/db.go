package database

import (
	"fmt"
	"log"

	"github.com/mcsans/finalProject3-kel2/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	user     = "postgres"
	password = "adminadmin"
	dbPort   = "5432"
	dbname   = "db_finalProject3-kel2"
	db       *gorm.DB
	err			 error
)

func StartDB() {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, dbPort)
	dsn := config
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("error connecting to database :", err)
	}

	fmt.Println("sukses koneksi ke database")
	db.Debug().AutoMigrate(models.User{}, models.Category{}, models.Task{})
}

func GetDB() *gorm.DB {
	return db
}