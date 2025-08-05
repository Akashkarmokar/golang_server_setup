package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/Akashkarmokar/golang_server_setup/internal/router"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout,
		&slog.HandlerOptions{AddSource: true, Level: slog.LevelInfo}),
	)
	logger.Info("Server starting on port 8080")
	r := router.New()
	if err := http.ListenAndServe(":8080", r); err != nil {
		logger.Error("Failed to start server:", "error", err)
	}
}
