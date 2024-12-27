package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"social-media-feed/internal/env"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"social-media-feed/internal/config"
	mwLogger "social-media-feed/internal/http/middleware/logger"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {

	env.MustLoad()
	cfg := config.MustLoad()

	fmt.Println(cfg)

	logger := setupLogger(cfg.Env)
	
	logger.Info("Start feed app")

	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	// r.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(mwLogger.New(logger))
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(cfg.Timeout))

	router.Get("/resources/*", func(w http.ResponseWriter, r *http.Request) {
		fs := http.StripPrefix("/resources/", http.FileServer(http.Dir("./resources")))
		fs.ServeHTTP(w, r)
	})
  
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fs := http.StripPrefix("/", http.FileServer(http.Dir("./resources")))
		fs.ServeHTTP(w, r)
	})

	address := fmt.Sprintf("%s:%s", cfg.HTTPServer.Host, cfg.HTTPServer.Port)

	http.ListenAndServe(address, router)
}

func setupLogger(env string) *slog.Logger {
	curr_time := time.Now()
	err := createFolderIfNotExists("logs")
	if err != nil {
		panic(err)
	}
	f, err := os.Create(fmt.Sprintf("./logs/feed_%s.txt", curr_time.Format(time.RFC3339)))
	if err != nil {
		panic(err)
	}
	
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(f, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewTextHandler(f, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewTextHandler(f, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	default: // If env config is invalid, set prod settings by default due to security
		log = slog.New(
			slog.NewTextHandler(f, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}

func createFolderIfNotExists(path string) error {
	err := os.MkdirAll(path, 0750)
	if os.IsExist(err) {
		return nil
	}
    return err
}