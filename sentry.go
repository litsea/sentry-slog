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

	if lc.sentryHub == nil {
		lc.sentryHub = sentry.NewHub(nil, sentry.NewScope())
	}

	client, err := sentry.NewClient(*o)
	if err != nil {
		return nil, fmt.Errorf("sentry.NewClient: %w", err)
	}

	lc.sentryHub.BindClient(client)

	return newLogHandler(lc.sentryHub, so, lc), nil
}
