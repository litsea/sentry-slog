package sentry

import (
	"context"
	"log/slog"

	"github.com/getsentry/sentry-go"
	slogsentry "github.com/samber/slog-sentry/v2"
)

type Option func(opt *sentry.ClientOptions)

func WithDSN(dsn string) Option {
	return func(opt *sentry.ClientOptions) {
		opt.Dsn = dsn
	}
}

func WithEnvironment(env string) Option {
	return func(opt *sentry.ClientOptions) {
		opt.Environment = env
	}
}

func WithAttachStacktrace(v bool) Option {
	return func(opt *sentry.ClientOptions) {
		opt.AttachStacktrace = v
	}
}

func WithRelease(rev string) Option {
	return func(opt *sentry.ClientOptions) {
		opt.Release = rev
	}
}

func WithDebug(v bool) Option {
	return func(opt *sentry.ClientOptions) {
		opt.Debug = v
	}
}

type LogOption func(lc *logConfig, opt *slogsentry.Option)

func LogWithLevel(l slog.Level) LogOption {
	return func(_ *logConfig, opt *slogsentry.Option) {
		opt.Level = l
	}
}

func LogWithAddSource(v bool) LogOption {
	return func(_ *logConfig, opt *slogsentry.Option) {
		opt.AddSource = v
	}
}

func LogWithAttrFromContext(fns ...func(ctx context.Context) []slog.Attr) LogOption {
	return func(_ *logConfig, opt *slogsentry.Option) {
		opt.AttrFromContext = fns
	}
}

func LogWithExtraAttrs(attrs map[string]string) LogOption {
	return func(lc *logConfig, _ *slogsentry.Option) {
		lc.extraAttrs = attrs
	}
}
