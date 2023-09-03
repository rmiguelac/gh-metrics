package metrics

import (
	"log"

	"github.com/rmiguelac/gh-metrics/pkg/workflows"
)

type Metrics struct {
	TotalJobs             int64
	TotalJobFailures      int64
	TotalWorkflows        int64
	TotalWorkflowFailures int64
}

/* Gather all metrics */
func GetMetrics(r *workflows.MyWorkflowRuns) *Metrics {

	log.Printf("Analysing %d workflows", r.Count)
	m := GetFailures(r)
	return m

}

/* Should give insights on how much we're failing builds */
func GetFailures(r *workflows.MyWorkflowRuns) *Metrics {

	log.Println("Gathering Job failures")

	total_job_failures := 0
	total_jobs := 0
	total_w_failures := 0
	total_w := 0
	for _, w := range r.Workflows {
		if w.Status == "failure" {
			total_w_failures += 1
		}
		for _, j := range w.MyJobs {
			if j.Status == "failure" {
				total_job_failures += 1
			}
			total_jobs += 1
		}
		total_w += 1
	}

	log.Printf("Total Jobs: %d", total_jobs)
	log.Printf("Total Failures: %d", total_job_failures)

	return &Metrics{TotalJobs: int64(total_jobs), TotalJobFailures: int64(total_job_failures), TotalWorkflows: int64(total_w), TotalWorkflowFailures: int64(total_w_failures)}
}

/* Get queue times of every job */
func GetQueueTimes() {}

/*
	Should gite insights regarding runners capacity and whether we should

increase its count
*/
func GetAverageQueueTime() {}

/*
	Should give insights on what steps could be improved regarding

times - either by runner resources or something else
*/
func GetLongestJobStep() {}

/* Get how long jobs are running */
func GetDurations() {}

/*
Should give insight on how much time it takes to get a feedback from a change
*/
func GetAverageDutarion() {}
