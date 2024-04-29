package db

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
	//if os.Getenv("GO_ENV") == "dev" {
	//err := godotenv.Load(fmt.Sprintf(".env.%s", os.Getenv("GO_ENV")))

	//ローカルの場合
	err := godotenv.Load(fmt.Sprintf(".env.dev"))

	//cloud runの場合
	//err := godotenv.Load(fmt.Sprintf(".env.prd"))

	if err != nil {
		log.Print(err)
	}
	//}
	//fmt.Printf("env getting")

	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PW"), os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_DB"))

	//fmt.Printf(url) //postgres://hato:hato72@localhost:5434/hato

	db, err := gorm.Open(postgres.Open(url), &gorm.Config{}) //データベースに接続 空の構造体をわたしている
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Connceted")
	return db
}

func CloseDB(db *gorm.DB) {
	sqlDB, _ := db.DB()
	if err := sqlDB.Close(); err != nil {
		log.Fatalln(err)
	}
}
