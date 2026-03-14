package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "subscriptions/docs"

	"subscriptions/internal/handler"
	"subscriptions/internal/repository"
	"subscriptions/internal/service"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Println(".env not found")
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPass, dbName,
	)

	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Database connected")

	repo := repository.NewSubscriptionRepository(db)
	service := service.NewSubscriptionService(repo)
	handler := handler.NewSubscriptionHandler(service)

	r := gin.Default()

	r.POST("/subscriptions", handler.Create)
	r.GET("/subscriptions", handler.GetAll)
	r.GET("/subscriptions/:id", handler.GetByID)
	r.PUT("/subscriptions/:id", handler.Update)
	r.DELETE("/subscriptions/:id", handler.Delete)
	r.GET("/subscriptions/total", handler.GetTotal)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":" + os.Getenv("SERVER_PORT"))
}
