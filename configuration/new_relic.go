package configuration

import (
	"fmt"

	"github.com/RianAsmara/personal-finance-advisor-api/exception"
	"github.com/newrelic/go-agent/v3/newrelic"
)

func NewRelicConfig(config Config) *newrelic.Application {
	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName(config.Get("APP_NAME")),
		newrelic.ConfigLicense(config.Get("NEW_RELIC_LICENSE")),
		newrelic.ConfigCodeLevelMetricsEnabled(true),
		newrelic.ConfigAppLogForwardingEnabled(true),
	)
	if nil != err {
		fmt.Printf("New Relic initialization failed: %v", err)
		exception.PanicLogging(err)
	}

	return app
}
