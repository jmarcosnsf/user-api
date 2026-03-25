package main

import (
	"log/slog"
	"net/http"
	"time"
	"user-api/internal/user"
)

func main() {
	if err := run(); err != nil {
		slog.Error("failed to execute code", "error", err)
		return
	}
	slog.Info("all systems offline")
}

func run() error {
	repo := user.NewRepository()
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