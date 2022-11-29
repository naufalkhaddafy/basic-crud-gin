package main

import (
	// book adalah directory root project go yang kita buat
	"learn-go/models" // memanggil package models pada directory models
	"learn-go/routes"
)

func main() {

    db := models.SetupDB()
    db.AutoMigrate(&models.Task{},&models.User{})


    r := routes.SetupRoutes(db)
    r.Run()
}