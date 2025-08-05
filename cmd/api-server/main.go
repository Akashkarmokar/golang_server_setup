package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/Akashkarmokar/golang_server_setup/internal/logger"
	"github.com/Akashkarmokar/golang_server_setup/internal/router"
)

func main() {
	log := slog.New(slog.NewJSONHandler(os.Stdout,
		&slog.HandlerOptions{AddSource: true, Level: slog.LevelInfo}),
	)

	r := router.New()
	wrappedRouter := logger.AddLoggerMid(log, logger.LoggerMid(r))
	log.Info("Server starting on port 8080")
	if err := http.ListenAndServe(":8080", wrappedRouter); err != nil {
		log.Error("Failed to start server:", "error", err)
	}
}
