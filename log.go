package sentry

import (
	"log/slog"

	"github.com/getsentry/sentry-go"
	slogsentry "github.com/samber/slog-sentry/v2"
)

var defaultLogExtraAttrs = map[string]string{
	"log.handler": "litsea.sentry",
}

type logConfig struct {
	extraAttrs map[string]string
}

func NewLogHandler(hub *sentry.Hub, opts ...LogOption) slog.Handler {
	o := &slogsentry.Option{
		Hub:       hub,
		Level:     slog.LevelError,
		AddSource: true,
	}

	lc := &logConfig{
		extraAttrs: defaultLogExtraAttrs,
	}

	for _, opt := range opts {
		opt(lc, o)
	}

	attrs := make([]slog.Attr, 0, len(lc.extraAttrs))
	for k, v := range lc.extraAttrs {
		attrs = append(attrs, slog.String(k, v))
	}

	return o.NewSentryHandler().WithAttrs(attrs)
}
