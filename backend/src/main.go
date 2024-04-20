package main

// import (
// 	"net/http"

// 	"github.com/labstack/echo/v4"
// )

// func main() {
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
)

func main() {
	db := db.NewDB()
	userValidator := validator.NewUserValidator()
	//taskValidator := validator.NewTaskValidator()
	userRepository := repository.NewUserRepository(db) //各コンストラクタ起動
	//taskRepository := repository.NewTaskRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository, userValidator)
	//taskUsecase := usecase.NewTaskUsecase(taskRepository, taskValidator)
	userController := controller.NewUserController(userUsecase)
	//taskController := controller.NewTaskController(taskUsecase)
	//e := router.NewRouter(userController, taskController)
	e := router.NewRouter(userController)
	e.Logger.Fatal(e.Start(":8080")) //サーバー起動
	//docker,pgAdminを起動->docker compose up -> bashでGO_ENV=dev go run backend/src/main.go
}
