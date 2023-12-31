package workflows

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/google/go-github/v54/github"
	"github.com/rmiguelac/gh-metrics/pkg/configuration"
	"github.com/rmiguelac/gh-metrics/pkg/utils"
	"golang.org/x/oauth2"
)

type MyWorkflowRun struct {
	Duration time.Duration
	MyJobs   []MyJob
	Workflow *github.Workflow
	Status   string
}

type MyWorkflowRuns struct {
	Count     int64
	Workflows []MyWorkflowRun
}

type MyJob struct {
	Duration   time.Duration
	QueuedTime time.Duration
	Status     string
	Steps      []*github.TaskStep
}

type MyJobs struct {
	Count int64
	Jobs  []MyJob
}

func GetMyWorkflows(c *configuration.Configuration) (*MyWorkflowRuns, error) {
	/* Do the Auth with oauth2 as recommended */
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: os.Getenv("GH_API_TOKEN")})
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	trange, err := utils.GetTimeRange(c)
	if err != nil {
		log.Printf("Unable to get expected time range")
	}

	w_opts := github.ListWorkflowRunsOptions{
		Created:     trange,
		ListOptions: github.ListOptions{PerPage: 10},
	}

	r := MyWorkflowRuns{}

	for {
		/* List all workflow runs */
		workflowRuns, resp, err := client.Actions.ListRepositoryWorkflowRuns(ctx, c.Organization, c.Repository, &w_opts)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		for _, w := range workflowRuns.WorkflowRuns {
			myw := MyWorkflowRun{
				Status: *w.Conclusion,
			}

			log.Printf("The workflowRun ID is: %d", w.ID)
			rid := *w.ID

			/* Get all jobs from current workflow */
			j_opts := github.ListWorkflowJobsOptions{
				ListOptions: github.ListOptions{PerPage: 10},
			}
			for {
				jobs, jresp, err := client.Actions.ListWorkflowJobs(ctx, c.Organization, c.Repository, rid, &j_opts)
				if err != nil {
					log.Fatal(err)
					return nil, err
				}

				for _, j := range jobs.Jobs {
					myj := MyJob{
						QueuedTime: j.StartedAt.Sub(j.CreatedAt.Time),
						Duration:   j.CompletedAt.Sub(j.StartedAt.Time),
						Steps:      j.Steps,
						Status:     *j.Conclusion,
					}
					log.Printf("The job id is: %d - Created at: %v - Started at: %v - Queue time: %v", *j.ID, *j.CreatedAt, *j.StartedAt, myj.QueuedTime)
					myw.MyJobs = append(myw.MyJobs, myj)
				}

				r.Workflows = append(r.Workflows, myw)
				r.Count += 1
				if jresp.NextPage == 0 {
					break
				}
				j_opts.Page = jresp.NextPage

			}
		}

		if resp.NextPage == 0 {
			break
		}
		w_opts.Page = resp.NextPage
	}

	return &r, nil
}
