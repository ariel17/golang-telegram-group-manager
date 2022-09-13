package config

import (
	"log"
	"os"

	"github.com/getsentry/sentry-go"
)

var (
	dsn string
)

func init() {
	dsn = os.Getenv("SENTRY_DSN")
	err := sentry.Init(sentry.ClientOptions{
		Dsn:              dsn,
		TracesSampleRate: 1.0,
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}
}