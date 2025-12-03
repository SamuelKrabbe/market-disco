package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/SamuelKrabbe/market-disco/api/config"
	"github.com/SamuelKrabbe/market-disco/api/internal/server"
	"github.com/SamuelKrabbe/market-disco/api/internal/storage/postgres"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)
	logger.Info("Starting api server")

	err := config.LoadConfig(".env")
	if err != nil {
		panic(err)
	}

	cfg, err := config.ParseConfig()
	if err != nil {
		panic(err)
	}

	db, err := postgres.NewPsqlDB(ctx, cfg)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	s := server.NewServer(ctx, cfg, db, *logger)
	if err := s.Run(ctx, s.Mount()); err != nil {
		slog.Error("server failed to start", "error", err)
		os.Exit(1)
	}
}
