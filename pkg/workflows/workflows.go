package workflows

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"

	constants "github.com/rmiguelac/gh-metrics/pkg/constants"
)

type APIClient struct {
	Headers http.Header
}

func NewClient() *APIClient {
	return &APIClient{
		Headers: *constants.GetHeaders(),
	}
}

/* Gather all workflow runs from a particular repository of a particular org
 */
func (a *APIClient) GetWorkflowRuns() {
	client := http.Client{}

	/* join github api URL with workflow runs endpoint */
	request_url, err := url.JoinPath(constants.GITHUB_API_URL, os.Getenv("GH_ORGANIZATION"), os.Getenv("GH_REPOSITORY"), constants.WORKFLOW_RUNS_API_URL)
	if err != nil {
		log.Fatal(err)
	}

	/* Prepare new request and add headers to it */
	req, err := http.NewRequest("GET", request_url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header = a.Headers

	res, err := client.Do(req)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := io.ReadAll(res.Body)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(responseData)

}

func (a *APIClient) GetJobRuns() {}

func (a *APIClient) GetJobQueuedTimes() {}

func (a *APIClient) GetJobFailures() {}
