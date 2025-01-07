package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"social-media-feed/internal/env"
	"social-media-feed/internal/fake"
	"social-media-feed/internal/repository"

	_ "github.com/lib/pq"
	ssov1 "github.com/username/protos/gen/go/sso"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

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

	defer db.Close()

	grpc_address := fmt.Sprintf("%s:%s", cfg.GRPCServer.Host, cfg.GRPCServer.Port)
	conn, err := grpc.NewClient(grpc_address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(fmt.Sprintf("Not connected with sso server: %s", err))
	}

	defer conn.Close()

	authClient := ssov1.NewAuthClient(conn)

	resp, err := authClient.IsAdmin(context.Background(), &ssov1.IsAdminRequest{UserId: fake.AdminId})
	if err != nil {
		panic(fmt.Sprintf("Not get reponse from sso server: %s", err))
	}

	logger.Info(fmt.Sprintf("Get IsAdmin response: %s", resp))

	repo := repository.NewRepository(db)
	handler := handler.NewHandler(logger, repo, &authClient)

	router, err := handler.InitRoutes(cfg)
	if err != nil {
		panic(fmt.Sprintf("Not init router: %s", err))
	}

	http_address := fmt.Sprintf("%s:%s", cfg.HTTPServer.Host, cfg.HTTPServer.Port)

	http.ListenAndServe(http_address, router)
}
