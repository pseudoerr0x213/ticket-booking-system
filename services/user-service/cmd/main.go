package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/pseudoerr0x213/ticket-booking-system/user-service/internal/config"
	"github.com/pseudoerr0x213/ticket-booking-system/user-service/internal/handler"
	"github.com/pseudoerr0x213/ticket-booking-system/user-service/internal/repository"
	"github.com/pseudoerr0x213/ticket-booking-system/user-service/internal/repository/postgres"
	"github.com/pseudoerr0x213/ticket-booking-system/user-service/internal/service"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

func main() {
	db, err := postgres.NewPostgresDB(config.PostgresConfig{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})

	if err != nil {
		return
	}
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	repo := repository.NewRepository(db, rdb)

	auth := service.NewAuthService(repo)

	src := service.NewServices(repo)
	mw := handler.NewMiddleware(*auth)

	handlers := handler.NewHandler(src, mw)

	router := handlers.SetupRoutes()
	server := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	log.Println("Server starting on :8080")
	log.Fatal(server.ListenAndServe())
}
