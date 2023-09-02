package main

import (
	"github.com/rmiguelac/gh-metrics/pkg/workflows"
)

func main() {
	apiclient := workflows.NewClient()
	apiclient.GetWorkflowRuns()
}
