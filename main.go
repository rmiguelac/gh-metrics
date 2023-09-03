package main

import (
	"log"

	"github.com/rmiguelac/gh-metrics/pkg/metrics"
	"github.com/rmiguelac/gh-metrics/pkg/workflows"
)

func main() {

	myw, err := workflows.GetMyWorkflows()
	if err != nil {
		log.Fatal(err)
	}

	m := metrics.GetMetrics(myw)
	metrics.RenderMetricsHTML(m)

}
