package handler

import (
	"log/slog"
	"net/http"
	"social-media-feed/internal/config"
	"social-media-feed/internal/repository"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	mwLogger "social-media-feed/internal/http/middleware/logger"
)

type Handler struct {
	logger *slog.Logger
	repo *repository.Repository
}

func NewHandler(logger *slog.Logger, repo *repository.Repository) *Handler {
	return &Handler{logger: logger, repo: repo}
}

func (h *Handler) InitRoutes(cfg *config.Config) (*chi.Mux, error) {

	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	// r.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(mwLogger.New(h.logger))
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(cfg.Timeout))
  
	router.Get("/", h.mainPage)
	
	router.Get("/resources/*", func(w http.ResponseWriter, r *http.Request) {
		fs := http.StripPrefix("/resources/", http.FileServer(http.Dir("./resources")))
		fs.ServeHTTP(w, r)
	})

	// Реализуйте тут свои обработчики
	router.Get("/createPost", h.createItem) 
	router.Post("/createPost",h.postFormCreateItem)

	return router, nil
}