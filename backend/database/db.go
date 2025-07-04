package database

import (
	"fmt"
	"log"

	"github.com/Gylmynnn/go-gin-flutter-bloc/config"
	"github.com/Gylmynnn/go-gin-flutter-bloc/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dbUser := config.GetEnv("DB_USER", "postgres")
	dbPass := config.GetEnv("DB_PASS", "")
	dbHost := config.GetEnv("DB_HOST", "localhost")
	dbPort := config.GetEnv("DB_POST", "5432")
	dbName := config.GetEnv("DB_NAME", "")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Makassar",
		dbHost, dbUser, dbPass, dbName, dbPort)

	var err error

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("pringatan : koneksi ke database gagal")
	}
	fmt.Println("info : berhasil terkoneksi ke database")

	err = DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("pringatan : migrasi model ke database gagal")
	}
	fmt.Println("info : berhasil migrasi model ke database")
}
