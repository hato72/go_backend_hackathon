package main

// import (
// 	"net/http"

// 	"database/sql"
// 	"fmt"

// 	"github.com/labstack/echo/v4"
// 	_ "github.com/lib/pq"
// )

// func main() {
// 	db, err := sql.Open("postgres", "user=hato password=hato72 dbname=hato host=db sslmode=disable")
// 	if err != nil {
// 		fmt.Println("DB connection error:", err)
// 		return
// 	}
// 	defer db.Close()

// 	err = db.Ping()
// 	if err != nil {
// 		fmt.Println("DB ping error:", err)
// 		return
// 	}

// 	// Query current date and time from the database
// 	var currentTime string
// 	err = db.QueryRow("SELECT NOW()").Scan(&currentTime)
// 	if err != nil {
// 		fmt.Println("DB query error:", err)
// 		return
// 	}

// 	fmt.Println("Current date and time:", currentTime)
// 	e := echo.New()
// 	e.GET("/", func(c echo.Context) error {
// 		return c.String(http.StatusOK, "Hello, World!")

// 	})
// 	e.Logger.Fatal(e.Start(":8080"))
// 	//fmt.Println("hello world hallo")
// }

import (
	"backend/src/controller"
	"backend/src/db"
	"backend/src/repository"
	"backend/src/router"
	"backend/src/usecase"
	"backend/src/validator"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	db := db.NewDB()
	userValidator := validator.NewUserValidator()
	cuisineValidator := validator.NewCuisineValidator()
	//taskValidator := validator.NewTaskValidator()

	userRepository := repository.NewUserRepository(db) //各コンストラクタ起動
	cuisineRepository := repository.NewCuisineRepository(db)
	//taskRepository := repository.NewTaskRepository(db)

	userUsecase := usecase.NewUserUsecase(userRepository, userValidator)
	cuisineUsecase := usecase.NewCuisineUsecase(cuisineRepository, cuisineValidator)
	//taskUsecase := usecase.NewTaskUsecase(taskRepository, taskValidator)

	userController := controller.NewUserController(userUsecase)
	cuisineController := controller.NewCuisineController(cuisineUsecase)
	//taskController := controller.NewTaskController(taskUsecase)

	//e := router.NewRouter(userController, taskController)
	e := router.NewRouter(userController, cuisineController)
	//e := router.NewRouter(userController)
	e.Logger.Fatal(e.Start(":8080")) //サーバー起動
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "hello world")
	})

	//docker,pgAdminを起動->docker compose up -> bashでGO_ENV=dev go run src/main.go
	//go run src/main.go
}
