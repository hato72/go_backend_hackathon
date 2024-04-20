package main

import (
	"backend/src/db"
	"backend/src/model"
	"fmt"
)

func main() {
	dbConn := db.NewDB()
	defer fmt.Println("Successfully Migrated")
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(&model.User{})
}

//GO_ENV=dev go run backend/src/migrate/migrate.go
