package app

import (
	"log/slog"
	grpcapp "ssov1/internal/app/grpc"
	"time"
)

type App struct {
	GRPCSv *grpcapp.App
}

func New(
	log *slog.Logger,
	grpcPort int,
	storagePath string,
	tokenTTL time.Duration) *App {
	// TODO: инициализировать хранилище
	// TODO: инициализировать сервис авторизации

	grpcApp := grpcapp.New(log, grpcPort)
	return &App{
		GRPCSv: grpcApp,
	}
}
