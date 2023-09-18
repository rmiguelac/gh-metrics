package main

import (
	"log"

	"github.com/rmiguelac/gh-metrics/pkg/configuration"
	"github.com/rmiguelac/gh-metrics/pkg/metrics"
	"github.com/rmiguelac/gh-metrics/pkg/workflows"
)

func main() {

	c := configuration.New()
	myw, err := workflows.GetMyWorkflows(c)
	if err != nil {
		log.Fatal(err)
	}

	m := metrics.GetMetrics(myw)
	metrics.RenderMetricsHTML(c, m)
	metrics.ToInfluxDB(c, m)

}
