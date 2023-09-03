package metrics

import (
	"io"
	"log"
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
)

/* Given the dataset, plot a piechart */
func PlotPieChart(m *Metrics) *charts.Pie {

	/* Get the data from Metrics */
	items := make([]opts.PieData, 0)
	items = append(items, opts.PieData{Name: "Total Workflows", Value: m.TotalWorkflows})
	items = append(items, opts.PieData{Name: "Total Workflow Failures", Value: m.TotalWorkflowFailures})
	pie := charts.NewPie()
	pie.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "Workflow Failure Rate"}),
	)
	pie.AddSeries("Workflow", items)

	return pie
}

/* Gather all metrics and return an HTML page */
func RenderMetricsHTML(m *Metrics) {

	page := components.NewPage()
	pie := PlotPieChart(m)
	page.AddCharts(pie)

	f, err := os.Create("metrics.html")
	if err != nil {
		log.Printf("Unable to craete html w/ metrics %s", err)
	}

	page.Render(io.MultiWriter(f))

}
