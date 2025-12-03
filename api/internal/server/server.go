package server

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/SamuelKrabbe/market-disco/api/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Server struct
type Server struct {
	ctx    context.Context
	cfg    *config.Config
	db     *pgxpool.Pool
	logger slog.Logger
}

func NewServer(ctx context.Context, cfg *config.Config, db *pgxpool.Pool, logger slog.Logger) *Server {
	return &Server{
		ctx:    ctx,
		cfg:    cfg,
		db:     db,
		logger: logger,
	}
}

func (s *Server) Mount() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("all good"))
	})

	// mux.HandleFunc("POST /task/", s.createTaskHandler)
	// mux.HandleFunc("GET /task/", s.getAllTasksHandler)
	// mux.HandleFunc("DELETE /task/", s.deleteAllTasksHandler)
	// mux.HandleFunc("GET /task/{id}/", s.getTaskHandler)
	// mux.HandleFunc("DELETE /task/{id}/", s.deleteTaskHandler)
	// mux.HandleFunc("GET /tag/{tag}/", s.tagHandler)
	// mux.HandleFunc("GET /due/{year}/{month}/{day}/", s.dueHandler)

	return mux
}

func (s *Server) Run(ctx context.Context, h http.Handler) error {
	srv := &http.Server{
		Addr:    ":" + s.cfg.Server.Port,
		Handler: h,
	}

	// Run the server in the background
	go func() {
		s.logger.Info("server listening", "address", srv.Addr)

		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			s.logger.Error("server error", "error", err)
		}
	}()

	// Wait for CTRL+C / kill signal
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	select {
	case <-stop:
		s.logger.Warn("received OS shutdown signal…")
	case <-ctx.Done():
		s.logger.Warn("context canceled…")
	}

	// graceful shutdown
	shutdownCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		s.logger.Error("graceful shutdown failed, forcing close", "error", err)

		// fallback hard kill
		if cerr := srv.Close(); cerr != nil {
			s.logger.Error("force close failed", "error", cerr)
		}
		return err
	}

	s.logger.Info("Server Stopped Gracefully")
	return nil
}
