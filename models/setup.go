package models

import (
  "fmt"
  "os"
  "github.com/joho/godotenv"
  "gorm.io/driver/mysql"
  "gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDataBase() {
	err := godotenv.Load(".env")

	if err != nil {
		panic("Error loading .env file")
	}
	mysqlUserName := os.Getenv("MYSQL_USERNAME")
	mysqlPassword := os.Getenv("MYSQL_PASSWORD")
	databaseName := os.Getenv("DATABASE_NAME")
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", mysqlUserName, mysqlPassword, databaseName)
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

  if err != nil {
    panic("failed to connect database")
  }

  database.AutoMigrate(&Course{}, &CourseContent{})

  DB = database
}
