package main

import (
	"fmt"
	"log"

	"github.com/Gylmynnn/go-gin-flutter-bloc/config"
	"github.com/Gylmynnn/go-gin-flutter-bloc/database"
	"github.com/Gylmynnn/go-gin-flutter-bloc/router"
)

func main() {
	if err := config.LoadEnv(); err != nil {
		log.Println("peringatan : file .env tidak dapat ditemukan")
	}
	log.Println("info : file .env berhasil dimuat")
	if err := database.InitDB(); err != nil {
		log.Println("peringatan : gagal terkoneksi dengan database", err.Error())
	}

	r := router.InitRoutes()
	if err := r.Run(":" + config.GetEnv("APP_PORT", "3000")); err != nil {
		log.Fatal("peringatan : gagal menjalankan server")
	}

	fmt.Println("server berjalan pada port : 3000")
}
