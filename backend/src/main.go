package main

import (
	"backend/src/controller"
	"backend/src/db"
	"backend/src/model"
	"backend/src/repository"
	"backend/src/router"
	"backend/src/usecase"
	"backend/src/validator"
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	//migrate

	//ローカル
	dbConn := db.NewDB()
	defer fmt.Println("Successfully Migrated")
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(&model.User{}, &model.Cuisine{})

	db := db.NewDB()

	//本番環境
	// dbConn := db.NewPrdDB()
	// defer fmt.Println("Successfully Migrated")
	// defer db.ClosePrdDB(dbConn)
	// dbConn.AutoMigrate(&model.User{}, &model.Cuisine{})

	// db := db.NewPrdDB()

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

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	e.Logger.Fatal(e.Start(":" + port)) //サーバー起動
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "hello")
	})

	//dockerを起動->docker compose build -> docker compose up // -> docker compose run --rm backend sh=> go run src/main.go
	// docker-compose down -v --rmi local
}
