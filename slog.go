package sentry

import (
	"log/slog"

	"github.com/getsentry/sentry-go"
	slogsentry "github.com/samber/slog-sentry/v2"
)

var defaultLogExtraAttrs = map[string]string{
	"log.handler": "litsea.sentry-slog",
}

type logConfig struct {
	extraAttrs map[string]string
}

func newLogHandler(hub *sentry.Hub, so *slogsentry.Option, lc *logConfig) slog.Handler {
	so.Hub = hub

	attrs := make([]slog.Attr, 0, len(lc.extraAttrs))
	for k, v := range lc.extraAttrs {
		attrs = append(attrs, slog.String(k, v))
	}

	return so.NewSentryHandler().WithAttrs(attrs)
}
