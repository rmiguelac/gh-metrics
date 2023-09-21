package metrics

import (
	"io"
	"log"
	"os"
	"regexp"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/rmiguelac/gh-metrics/pkg/configuration"
)

/* Receive a Map and transform it into PieData */
func PreparePieChartData(d *map[string]any, c *configuration.Configuration) []opts.PieData {

	success_color := "rgba(" + c.Report.Colors.Success + ")"
	failure_color := "rgba(" + c.Report.Colors.Failure + ")"
	items := make([]opts.PieData, 0)
	/* Case insensiteve search for failure */
	reg := regexp.MustCompile(`(?i)Fail`)
	for k, v := range *d {
		i_color := success_color
		if reg.MatchString(k) {
			i_color = failure_color
		}
		items = append(items, opts.PieData{Name: k, Value: v, ItemStyle: &opts.ItemStyle{Color: i_color}})
	}
	return items
}

/* Given the dataset, plot a piechart */
func PieChart(d []opts.PieData) *charts.Pie {

	/* Get the data from Metrics */

	pie := charts.NewPie()

	pie.AddSeries("", d).SetSeriesOptions(
		charts.WithLabelOpts(
			opts.Label{
				Show:      true,
				Formatter: "{b}: {c}",
			},
		),
	)

	return pie
}

func PrepareWorkflowPieChart(m *Metrics, c *configuration.Configuration) *charts.Pie {

	data := map[string]any{
		"Total Workflows":         m.TotalWorkflows,
		"Total Workflow Failures": m.TotalWorkflowFailures,
	}
	pdata := PreparePieChartData(&data, c)
	pie := PieChart(pdata)
	pie.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "Workflow Failure Rate"}),
	)
	return pie

}

func PrepareJobsPieChart(m *Metrics, c *configuration.Configuration) *charts.Pie {

	data := map[string]any{
		"Total Jobs":         m.TotalJobs,
		"Total Job Failures": m.TotalJobFailures,
	}
	pdata := PreparePieChartData(&data, c)
	pie := PieChart(pdata)
	pie.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "Jobs Failure Rate"}),
	)
	return pie

}

/* Gather all metrics and return an HTML page */
func RenderMetricsHTML(c *configuration.Configuration, m *Metrics) {

	page := components.NewPage()
	wpie := PrepareWorkflowPieChart(m, c)
	page.AddCharts(wpie)
	jpie := PrepareJobsPieChart(m, c)
	page.AddCharts(jpie)

	f, err := os.Create("metrics.html")
	if err != nil {
		log.Printf("Unable to craete html w/ metrics %s", err)
	}

	page.Render(io.MultiWriter(f))

}
