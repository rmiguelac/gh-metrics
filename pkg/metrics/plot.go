package metrics

import (
	"io"
	"log"
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
)

/* Receive a Map and transform it into PieData */
func PreparePieChartData(d *map[string]any) []opts.PieData {

	items := make([]opts.PieData, 0)
	for k, v := range *d {
		items = append(items, opts.PieData{Name: k, Value: v})
	}
	return items
}

/* Given the dataset, plot a piechart */
func PieChart(d []opts.PieData) *charts.Pie {

	/* Get the data from Metrics */

	pie := charts.NewPie()
	pie.AddSeries("", d)

	return pie
}

func PrepareWorkflowPieChart(m *Metrics) *charts.Pie {

	data := map[string]any{
		"Total Workflows":         m.TotalWorkflows,
		"Total Workflow Failures": m.TotalWorkflowFailures,
	}
	pdata := PreparePieChartData(&data)
	pie := PieChart(pdata)
	pie.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "Workflow Failure Rate"}),
	)
	return pie

}

func PrepareJobsPieChart(m *Metrics) *charts.Pie {

	data := map[string]any{
		"Total Jobs":         m.TotalJobs,
		"Total Job Failures": m.TotalJobFailures,
	}
	pdata := PreparePieChartData(&data)
	pie := PieChart(pdata)
	pie.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "Jobs Failure Rate"}),
	)
	return pie

}

/* Gather all metrics and return an HTML page */
func RenderMetricsHTML(m *Metrics) {

	page := components.NewPage()
	wpie := PrepareWorkflowPieChart(m)
	page.AddCharts(wpie)
	jpie := PrepareJobsPieChart(m)
	page.AddCharts(jpie)

	f, err := os.Create("metrics.html")
	if err != nil {
		log.Printf("Unable to craete html w/ metrics %s", err)
	}

	page.Render(io.MultiWriter(f))

}
