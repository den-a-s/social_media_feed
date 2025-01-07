package handler

import (
	"log/slog"
	"net/http"
	"social-media-feed/internal/config"
	"social-media-feed/internal/repository"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	mwLogger "social-media-feed/internal/http/middleware/logger"

	ssov1 "github.com/username/protos/gen/go/sso"
)

type Handler struct {
	logger *slog.Logger
	repo *repository.Repository
	authClient *ssov1.AuthClient
}

func NewHandler(logger *slog.Logger, repo *repository.Repository, authClient *ssov1.AuthClient) *Handler {
	return &Handler{logger: logger, repo: repo, authClient: authClient}
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

	return router, nil
}