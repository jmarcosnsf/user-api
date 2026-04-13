package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"time"
	"user-api/internal/user"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func main() {
	if err := run(); err != nil {
		slog.Error("failed to execute code", "error", err)
		return
	}
	slog.Info("all systems offline")
}

func run() error {
	ctx := context.Background()
	godotenv.Load()

	pool, err := pgxpool.New(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		return err
	}

	repo := user.NewRepository(pool)
	handler := user.NewHandler(repo)

	s := http.Server{
		ReadTimeout: 10 * time.Second,
		IdleTimeout: time.Minute,
		WriteTimeout: 10 * time.Second,
		Addr: ":8080",
		Handler: handler,
	}

	if err := s.ListenAndServe(); err != nil {
		return err 
	}

	return nil
}