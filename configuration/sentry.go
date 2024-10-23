package configuration

import (
	"fmt"
	"os"

	"github.com/getsentry/sentry-go"
)

func SentryConfig(config Config) {
	if err := sentry.Init(sentry.ClientOptions{
		Dsn:              config.Get("SENTRY_DSN"),
		EnableTracing:    true,
		TracesSampleRate: 1.0,
	}); err != nil {
		fmt.Printf("Sentry initialization failed: %v", err)
		os.Exit(1)
	}
}
