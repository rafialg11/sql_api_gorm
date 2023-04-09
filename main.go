package main

import (
	"os"
	"sql-api-gorm/database"
	"sql-api-gorm/routers"
)

func main() {
	database.StartDB()
	var PORT = os.Getenv("PORT")
	routers.StartServer().Run(PORT)
}
