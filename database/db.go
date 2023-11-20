package database

import (
	"fmt"
	"log"

	"github.com/mcsans/finalProject3-kel2/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "roundhouse.proxy.rlwy.net"
	user     = "postgres"
	password = "Dd*afb5F*1gC66Df5caFEcdgCAC3DBDC"
	dbPort   = "15087"
	dbname   = "railway"
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