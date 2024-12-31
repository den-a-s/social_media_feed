package router

import (
	"fmt"
	"log/slog"
	"net/http"
	"social-media-feed/internal/config"
	"social-media-feed/internal/http/template"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	mwLogger "social-media-feed/internal/http/middleware/logger"
)

func InitRouter(cfg *config.Config,  logger *slog.Logger) (*chi.Mux, error) {

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
		main_html, err := template.GetFilledMainTemplate(logger)
		if err != nil {
			logger.Error(fmt.Sprintf("Не смогли получить сформированный шаблон: %s", err))
		}

		w.Write([]byte(main_html))
	})

	return router, nil
}