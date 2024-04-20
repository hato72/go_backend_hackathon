package main

import (
	"net/http"
	"database/sql"
	"github.com/labstack/echo/v4"
	"fmt"
	_ "github.com/lib/pq"
)

func main() {
    db, err := sql.Open("postgres", "user=hato password=hato72 dbname=hato host=db sslmode=disable")
    if err != nil {
        fmt.Println("DB connection error:", err)
        return
    }
    defer db.Close()

    err = db.Ping()
    if err != nil {
        fmt.Println("DB ping error:", err)
        return
    }

    // Query current date and time from the database
    var currentTime string
    err = db.QueryRow("SELECT NOW()").Scan(&currentTime)
    if err != nil {
        fmt.Println("DB query error:", err)
        return
    }

    fmt.Println("Current date and time:", currentTime)
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":8080"))
	//fmt.Println("hello world hallo")
}
