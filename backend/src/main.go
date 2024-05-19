package main

import (
	"backend/graph"
	"backend/src/controller"
	"backend/src/db"
	"backend/src/model"
	"backend/src/repository"
	"backend/src/router"
	"backend/src/usecase"
	"backend/src/validator"
	"fmt"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/hato72/go_backend_hackathon/backend/generated"
	"github.com/hato72/go_backend_hackathon/backend/graph"
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

	graphqlHandler := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{Resolvers: &graph.Resolver{DB: db}},
		),
	)
	playgroundHandler := playground.Handler("GraphQL", "/query")

	e.POST("/query", func(c echo.Context) error {
		graphqlHandler.ServeHTTP(c.Response(), c.Request())
		return nil
	})

	e.GET("/playground", func(c echo.Context) error {
		playgroundHandler.ServeHTTP(c.Response(), c.Request())
		return nil
	})

	e.Logger.Fatal(e.Start(":8080")) //サーバー起動
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "hello world")
	})

	//dockerを起動->docker compose build -> docker compose up // -> docker compose run --rm backend sh=> go run src/main.go
	// docker-compose down -v --rmi local
}
