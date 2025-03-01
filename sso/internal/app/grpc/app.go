package grpcapp

import (
	"fmt"
	"log/slog"
	"net"
	authrpc "sso/internal/grpc/auth"

	"google.golang.org/grpc"
)

type App struct {
	log        *slog.Logger
	gRPCServer *grpc.Server
	port       int
}

// Метод создает новый gRPC сервер
func New(log *slog.Logger, authService authrpc.Auth, port int) *App {
	gRPCServer := grpc.NewServer()
	authrpc.Register(gRPCServer, authService)
	return &App{
		log:        log,
		gRPCServer: gRPCServer,
		port:       port,
	}
}

// Запуск приложения
func (a *App) Run() error {
	const op = "grpcApp.Run"
	log := a.log.With(slog.String("op", op), slog.Int("port", a.port))

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	log.Info("grpc server is running", slog.String("addr:", l.Addr().String()))
	if err := a.gRPCServer.Serve(l); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

// Запуск приложения. Паникует в случае неудачи
func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

// Остановка приложения
func (a *App) Stop() {
	const op = "grpcapp.Stop"
	a.log.With(slog.String("op", op)).Info("stopping gRPC server", slog.Int("port", a.port))

	a.gRPCServer.GracefulStop()
}
