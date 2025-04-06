package sentry

import (
	"fmt"
	"log/slog"

	"github.com/getsentry/sentry-go"
	slogsentry "github.com/samber/slog-sentry/v2"
)

func NewHandler(opts ...Option) (slog.Handler, error) {
	o := &sentry.ClientOptions{
		AttachStacktrace: true,
	}

	so := &slogsentry.Option{
		Level:     slog.LevelError,
		AddSource: true,
	}

	lc := &logConfig{
		extraAttrs: defaultLogExtraAttrs,
	}

	for _, opt := range opts {
		opt(o, so, lc)
	}

	hub := sentry.NewHub(nil, sentry.NewScope())
	client, err := sentry.NewClient(*o)
	if err != nil {
		return nil, fmt.Errorf("sentry.NewClient: %w", err)
	}

	hub.BindClient(client)

	return newLogHandler(hub, so, lc), nil
}
