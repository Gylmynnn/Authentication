package router

import (
	"github.com/Gylmynnn/go-gin-flutter-bloc/controllers"
	"github.com/Gylmynnn/go-gin-flutter-bloc/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRoutes() *gin.Engine {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:  []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders: []string{"Content-Length"},
	}))
	r.POST("/api/register", controllers.Register)
	r.POST("/api/login", controllers.Login)
	r.GET("/api/users", middleware.AuthMiddleware(), controllers.FindAllUsers)
	r.POST("/api/users", middleware.AuthMiddleware(), controllers.CreateUser)
	r.GET("/api/users/:id", middleware.AuthMiddleware(), controllers.FindUserById)
	r.PUT("/api/users/:id", middleware.AuthMiddleware(), controllers.UpdateUser)
	r.DELETE("/api/users/:id", middleware.AuthMiddleware(), controllers.DeleteUser)
	return r
}
