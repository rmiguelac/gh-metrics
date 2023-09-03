package main

import (
	"context"
	"log"
	"os"

	"github.com/google/go-github/v54/github"
	"golang.org/x/oauth2"
)

func main() {

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: os.Getenv("GH_API_TOKEN")})
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	workflowRuns, _, err := client.Actions.ListRepositoryWorkflowRuns(ctx, os.Getenv("GH_ORGANIZATION"), os.Getenv("GH_REPOSITORY"), nil)
	if err != nil {
		log.Fatal(err)
	}

	/* For every workflow, get their job url
	 */
	for _, w := range workflowRuns.WorkflowRuns {
		log.Printf("The workflowRun ID is: %d", w.ID)
		rid := *w.ID
		jobs, _, err := client.Actions.ListWorkflowJobs(ctx, os.Getenv("GH_ORGANIZATION"), os.Getenv("GH_REPOSITORY"), rid, nil)
		if err != nil {
			log.Fatal(err)
		}
		for _, j := range jobs.Jobs {
			delta := j.StartedAt.Sub(j.CreatedAt.Time)
			log.Printf("The job id is: %d - Created at: %v - Started at: %v - Queue time: %v", *j.ID, *j.CreatedAt, *j.StartedAt, delta)
		}
	}
}
