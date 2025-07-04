package main

import (
	"fmt"
	"log"

	"github.com/Gylmynnn/go-gin-flutter-bloc/config"
	"github.com/Gylmynnn/go-gin-flutter-bloc/database"
	"github.com/Gylmynnn/go-gin-flutter-bloc/router"
)

func main() {
	config.LoadEnv()
	database.InitDB()

	r := router.InitRoutes()
	err := r.Run(":" + config.GetEnv("APP_PORT", "3000"))
	if err != nil {
		log.Fatal("peringatan : gagal menjalankan server")
	}
	fmt.Println("server berjalan pada port : 3000")
}
