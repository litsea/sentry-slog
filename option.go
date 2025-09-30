package sentry

import (
	"context"
	"log/slog"

	"github.com/getsentry/sentry-go"
	slogsentry "github.com/samber/slog-sentry/v2"
)

type Option func(opt *sentry.ClientOptions, _ *slogsentry.Option, _ *logConfig)

func WithDSN(dsn string) Option {
	return func(opt *sentry.ClientOptions, _ *slogsentry.Option, _ *logConfig) {
		opt.Dsn = dsn
	}
}

func WithEnvironment(env string) Option {
	return func(opt *sentry.ClientOptions, _ *slogsentry.Option, _ *logConfig) {
		opt.Environment = env
	}
}

func WithAttachStacktrace(v bool) Option {
	return func(opt *sentry.ClientOptions, _ *slogsentry.Option, _ *logConfig) {
		opt.AttachStacktrace = v
	}
}

func WithRelease(rev string) Option {
	return func(opt *sentry.ClientOptions, _ *slogsentry.Option, _ *logConfig) {
		opt.Release = rev
	}
}

func WithDebug(v bool) Option {
	return func(opt *sentry.ClientOptions, _ *slogsentry.Option, _ *logConfig) {
		opt.Debug = v
	}
}

func WithLogLevel(l slog.Level) Option {
	return func(_ *sentry.ClientOptions, opt *slogsentry.Option, _ *logConfig) {
		opt.Level = l
	}
}

func WithLogAddSource(v bool) Option {
	return func(_ *sentry.ClientOptions, opt *slogsentry.Option, _ *logConfig) {
		opt.AddSource = v
	}
}

func WithLogAttrFromContext(fns ...func(ctx context.Context) []slog.Attr) Option {
	return func(_ *sentry.ClientOptions, opt *slogsentry.Option, _ *logConfig) {
		if len(fns) > 0 && fns[0] != nil {
			opt.AttrFromContext = fns
		}
	}
}

func WithLogReplaceAttr(fns ...func(groups []string, a slog.Attr) slog.Attr) Option {
	return func(_ *sentry.ClientOptions, opt *slogsentry.Option, _ *logConfig) {
		if len(fns) > 0 && fns[0] != nil {
			opt.ReplaceAttr = fns[0]
		}
	}
}

func WithSentryHub(hub *sentry.Hub) Option {
	return func(_ *sentry.ClientOptions, _ *slogsentry.Option, lc *logConfig) {
		if hub != nil {
			lc.sentryHub = hub
		}
	}
}

func WithLogExtraAttrs(attrs map[string]string) Option {
	return func(_ *sentry.ClientOptions, _ *slogsentry.Option, lc *logConfig) {
		if len(attrs) > 0 {
			lc.extraAttrs = attrs
		}
	}
}
