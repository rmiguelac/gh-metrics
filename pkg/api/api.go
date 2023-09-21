package api

import (
	"github.com/rmiguelac/gh-metrics/pkg/configuration"
	"github.com/rmiguelac/gh-metrics/pkg/metrics"
)

type Collector struct {
	configuration *configuration.Configuration
	influxstore   *metrics.InfluxStore
}
