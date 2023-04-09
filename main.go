package main

import (
	"sql-api-gorm/database"
	"sql-api-gorm/routers"
)

func main() {
	database.StartDB()
	var PORT = ":4000"
	routers.StartServer().Run(PORT)
}
