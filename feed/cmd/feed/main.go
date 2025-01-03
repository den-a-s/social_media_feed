package main

import (
	"fmt"
	"net/http"
	"os"
	"social-media-feed/internal/env"
	"social-media-feed/internal/repository"

	_ "github.com/lib/pq"

	"social-media-feed/internal/config"
	"social-media-feed/internal/http/handler"
	"social-media-feed/internal/http/middleware/logger"
)

// TODO Врапнуть все паники так чтобы мой логгер записывал логи к себе
func main() {

	env.MustLoad()
	cfg := config.MustLoad()

	fmt.Println(cfg)

	logger := logger.SetupLogger(cfg.Env)
	
	logger.Info("Start feed app")

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     cfg.DB.Host,
		Port:     cfg.DB.Port,
		Username: cfg.DB.Username,
		DBName:   cfg.DB.DBName,
		SSLMode:  cfg.DB.SSLMode,
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		panic(fmt.Sprintf("Not connected to db: %s", err))
	}
	
	res, err := db.Queryx("SELECT * FROM schema_migrations")
	if err != nil {
		panic(fmt.Sprintf("DB select post err: %s", err))
	}

	logger.Info(fmt.Sprintf("res from select: %s", res))

	repo := repository.NewRepository(db)
	handler := handler.NewHandler(logger, repo)

	router, err := handler.InitRoutes(cfg)
	if err != nil {
		panic(fmt.Sprintf("Not init router: %s", err))
	}

	address := fmt.Sprintf("%s:%s", cfg.HTTPServer.Host, cfg.HTTPServer.Port)

	http.ListenAndServe(address, router)
}
