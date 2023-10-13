package main

import (
	"github.com/mcsans/finalProject3-kel2/database"
	"github.com/mcsans/finalProject3-kel2/router"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	r.Run(":8080")
}