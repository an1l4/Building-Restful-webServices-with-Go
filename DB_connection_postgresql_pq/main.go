package main

import (
	"db/connections/models"
	"log"
)

func main() {
	db, err := models.InitDB()
	if err != nil {
		log.Println(db)
	}
}
