package main

import (
	"CRUD-Golang/database"
	"CRUD-Golang/routers"
)

func main() {
	database.InitDatabase()
	router := routers.SetupRouter()
	router.Run(":8080")
}
