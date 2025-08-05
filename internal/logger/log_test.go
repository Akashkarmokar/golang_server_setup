package logger_test

import (
	"context"
	"log/slog"
	"os"
	"testing"

	"github.com/Akashkarmokar/golang_server_setup/internal/logger"
)

func Test_CtxWithLogger(t *testing.T) {
	testCases := []struct {
		name            string
		ctx             context.Context
		logger          *slog.Logger
		doesLoggerExist bool
	}{
		{
			name: "returns context without looger",
			ctx:  context.Background(),
		},
		{
			name: "return as it is",
			ctx: context.WithValue(
				context.Background(),
				logger.CtxKey{},
				slog.New(slog.NewJSONHandler(
					os.Stdout,
					&slog.HandlerOptions{AddSource: true},
				),
				),
			),
			doesLoggerExist: true,
		},
		{
			name: "Inject Logger",
			ctx:  context.Background(),
			logger: slog.New(slog.NewJSONHandler(
				os.Stdout,
				&slog.HandlerOptions{AddSource: true},
			)),
			doesLoggerExist: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := logger.CtxWithLogger(tc.ctx, tc.logger)
			_, ok := ctx.Value(logger.CtxKey{}).(*slog.Logger)
			if tc.doesLoggerExist != ok {
				t.Errorf("Expected : %v got %v", false, ok)
			}
		})
	}
}

func Test_FromContext(t *testing.T) {
	testCases := []struct {
		name     string
		ctx      context.Context
		expected bool
	}{
		{
			name: "logger exist",
			ctx: context.WithValue(
				context.Background(),
				logger.CtxKey{},
				slog.New(slog.NewJSONHandler(
					os.Stdout,
					&slog.HandlerOptions{AddSource: true},
				),
				),
			),
			expected: true,
		},
		{
			name:     "new logger returned",
			ctx:      context.Background(),
			expected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			logger := logger.FromContext(tc.ctx)

			if tc.expected && logger == nil {
				t.Errorf("Expected : %v, got %v", tc.expected, logger)
			}
		})
	}
}
