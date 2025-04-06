package sentry

import (
	"fmt"

	"github.com/getsentry/sentry-go"
)

func New(opts ...Option) (*sentry.Hub, error) {
	o := &sentry.ClientOptions{
		AttachStacktrace: true,
	}

	for _, opt := range opts {
		opt(o)
	}

	hub := sentry.NewHub(nil, sentry.NewScope())
	client, err := sentry.NewClient(*o)
	if err != nil {
		return nil, fmt.Errorf("sentry.NewClient: %w", err)
	}

	hub.BindClient(client)
	return hub, nil
}
