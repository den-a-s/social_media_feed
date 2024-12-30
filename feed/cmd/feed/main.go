package main

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"log/slog"
	"net/http"
	"os"
	"social-media-feed/internal/env"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

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
		main_html, err := GetFilledMainTemplate(logger)
		if err != nil {
			logger.Error(fmt.Sprintf("Не смогли получить сформированный шаблон: %s", err))
		}

		w.Write([]byte(main_html))
	})

	address := fmt.Sprintf("%s:%s", cfg.HTTPServer.Host, cfg.HTTPServer.Port)

	http.ListenAndServe(address, router)
}

func GetFilledMainTemplate(logger *slog.Logger) (string, error) {
	type Posts struct {
		Title string
		PostImg string
	}

	tmpl, err := os.ReadFile("./web/templates/main_tmpl.html")
	if err != nil {
		return "", errors.New(fmt.Sprintf("Not read file: %s", err))
	}

	t, err := template.New("webpage").Parse(string(tmpl))
	if err != nil {
		return "", errors.New(fmt.Sprintf("Bad create template: %s", err))
	}

	data := struct {
		Posts []Posts
	}{
		Posts: []Posts{
			Posts{
				Title: "Мафбосс 1",
				PostImg: "resources/posts_image/mathboss_1.jpg",
			},
			Posts{
				Title: "Мафбосс 2",
				PostImg: "resources/posts_image/mathboss_2.jpg",
			},
		},
	};

	buf := new(bytes.Buffer)

	err = t.Execute(buf, data)
	if err != nil {
		return "", errors.New(fmt.Sprintf("Bad parsing: %s", err))
	}

	return buf.String(), nil
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