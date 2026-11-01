package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"strings"

	"github.com/urfave/cli/v3"

	formacli "github.com/zhulik/forma/cmd/forma/internal/cli"
)

func main() {
	root := formacli.NewRoot()
	root.Before = configureLoggingBefore(root)

	if err := root.Run(context.Background(), os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func configureLoggingBefore(_ *cli.Command) func(context.Context, *cli.Command) (context.Context, error) {
	return func(ctx context.Context, cmd *cli.Command) (context.Context, error) {
		level := slog.LevelInfo
		if cmd.IsSet(formacli.FlagLogLevel) {
			var err error
			level, err = parseLogLevel(cmd.String(formacli.FlagLogLevel))
			if err != nil {
				return ctx, err
			}
		} else if v := os.Getenv(formacli.EnvLogLevel); v != "" {
			var err error
			level, err = parseLogLevel(v)
			if err != nil {
				return ctx, err
			}
		}

		handler := slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: level})
		slog.SetDefault(slog.New(handler))
		return ctx, nil
	}
}

func parseLogLevel(s string) (slog.Level, error) {
	switch strings.ToLower(strings.TrimSpace(s)) {
	case "debug":
		return slog.LevelDebug, nil
	case "info":
		return slog.LevelInfo, nil
	case "warn", "warning":
		return slog.LevelWarn, nil
	case "error":
		return slog.LevelError, nil
	default:
		return slog.LevelInfo, fmt.Errorf("invalid log level %q (want debug, info, warn, error)", s)
	}
}
