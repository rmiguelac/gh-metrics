package workflows

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/google/go-github/v54/github"
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

func GetMyWorkflows() (*MyWorkflowRuns, error) {
	/* Do the Auth with oauth2 as recommended */
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: os.Getenv("GH_API_TOKEN")})
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	workflowRuns, _, err := client.Actions.ListRepositoryWorkflowRuns(ctx, os.Getenv("GH_ORGANIZATION"), os.Getenv("GH_REPOSITORY"), nil)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	r := MyWorkflowRuns{}

	/* For every workflow, get it to a struct and get its jobs details
	 */
	for _, w := range workflowRuns.WorkflowRuns {
		myw := MyWorkflowRun{
			Status: *w.Status,
		}

		log.Printf("The workflowRun ID is: %d", w.ID)
		rid := *w.ID
		jobs, _, err := client.Actions.ListWorkflowJobs(ctx, os.Getenv("GH_ORGANIZATION"), os.Getenv("GH_REPOSITORY"), rid, nil)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}

		for _, j := range jobs.Jobs {
			myj := MyJob{
				QueuedTime: j.StartedAt.Sub(j.CreatedAt.Time),
				Duration:   j.CompletedAt.Sub(j.StartedAt.Time),
				Steps:      j.Steps,
			}
			log.Printf("The job id is: %d - Created at: %v - Started at: %v - Queue time: %v", *j.ID, *j.CreatedAt, *j.StartedAt, myj.QueuedTime)
			myw.MyJobs = append(myw.MyJobs, myj)
		}

		r.Workflows = append(r.Workflows, myw)
		r.Count += 1
	}

	return &r, nil
}
