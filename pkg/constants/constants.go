package contants

import (
	"net/http"
	"os"
)

/*
URLs should not end with forward slash '/'
*/

/* URL constants */
const GITHUB_API_URL = "https://api.github.com/repos"
const WORKFLOW_RUNS_API_URL = "/actions/runs"

/* HEADERs */
func GetHeaders() *http.Header {
	h := http.Header{}
	var HEADERS = map[string]string{
		"Accept":               "application/json",
		"X-GitHub-Api-Version": "2022-11-28",
		"Authorization":        "Bearer " + os.Getenv("GH_API_TOKEN"),
	}
	for k, v := range HEADERS {
		h.Add(k, v)
	}

	return &h
}
