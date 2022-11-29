package main

import (
	"learn-go/models"
	"learn-go/routes"
)

func main() {

    db := models.SetupDB()
    db.AutoMigrate(&models.Task{},&models.User{})


    r := routes.SetupRoutes(db)
    r.Run()
}