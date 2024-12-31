package main

import (
	"fmt"
	"net/http"
	"os"
	"social-media-feed/internal/env"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"social-media-feed/internal/config"
	"social-media-feed/internal/http/middleware/logger"
	"social-media-feed/internal/http/router"
)

func main() {

	env.MustLoad()
	cfg := config.MustLoad()

	fmt.Println(cfg)

	logger := logger.SetupLogger(cfg.Env)
	
	logger.Info("Start feed app")

	db_connected_string := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.DB.Host, cfg.DB.Port, cfg.DB.Username, cfg.DB.DBName, os.Getenv("DB_PASSWORD"), cfg.DB.SSLMode)

	fmt.Println(db_connected_string)

	db, err := sqlx.Connect("postgres", db_connected_string)
    if err != nil {
        panic(fmt.Sprintf("DB is not connected %s", err))
    }
	
	res, err := db.Queryx("SELECT * FROM schema_migrations")
	if err != nil {
		panic(fmt.Sprintf("DB select post err: %s", err))
	}

	fmt.Println("Пишем логи")
	logger.Info(fmt.Sprintf("res from select: %s", res))

	router, err := router.InitRouter(cfg, logger)
	if err != nil {
		panic(fmt.Sprintf("Not init router: %s", err))
	}

	address := fmt.Sprintf("%s:%s", cfg.HTTPServer.Host, cfg.HTTPServer.Port)

	http.ListenAndServe(address, router)
}
