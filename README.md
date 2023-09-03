# gh-metrics
Get metrics from github workflow runs

## Using

### Setup relevant environment variables

The following environment variables must be available:

| Environment Variable | Default | Required | Example    | Description |
| -------------------- | ------- | -------- | ---------- | ----------- |
| GH_ORGANIZATION      | ''      | true     | rmiguelac  | The owner/org of the repo to be analyzed |
| GH_REPOSITORY        | ''      | true     | gh-metrics | The repo to be analyzed |
| GH_API_TOKEN         | ''      | true     | GH-PAT     | The Personal Access Token |

To have them, simply do:

`export VAR=value`

### Run

Simply run it with `go run main.go` or build it with `go build .` then run it `./gh-metrics`


### Output

A `metrics.html` shouuld've been created in the same folder where the code is.

## Example Output

<p align='center'.>
  <img src="./static/metrics.png" width=50% height=50%>
</p>


## TODO

[] Support Pagination  
[] Support Query Workflows betwwen dates  
[] Show values in pie charts  