# gh-metrics
Get metrics from github workflow runs

## Using

### Setup relevant parameters

Parameters that can be configured in [config.yaml](config.yaml)

| Parameter | Default | Required | Example | Description |
| --------- | ------- | -------- | ------- | ----------- |
| organization | ''   | yes      | rmiguelac | The owner/org of the repo to be analyzed | 
| repository | ''     | yes      | gh-metrics | The repo to be analyzed |
| report.data.last | `days: 7` | no | `report.data.last.weeks: 2` | Number of `hours`, `days`, `weeks` or `months` to look for workflows. Just one option at the time |
| report.data.colors.failure | `report.colors.failure: "255, 61, 71, 0.7"` | no | `report.colors.failure: "255, 61, 71, 0.7"` | Color to be used for data points that represent failures. RGBA format |
| report.data.colors.success | `report.colors.success: "139, 193, 71, 0.7"` | no | `report.colors.success: "139, 193, 71, 0.7"` | Color to be used for data points that represent successs. RGBA format |

#### On Environment Variables and File Parameters

There are certain parameters that can be configured in multiple places. For example `organization` and `repository` parameters can be configured in [config.yaml](config.yaml) or
through the `GH_ORGANIZATION` and `GH_REPOSITORY` environment variables.  

The precedence order is as follows:
* Environment variables 
* Config file parameters

That means that if an environment variable is set for the same parameter in the config file, the environment variable takes precedence.

### Environment variables

| Environment Variable | Default | Required | Example    | Description |
| -------------------- | ------- | -------- | ---------- | ----------- |
| GH_ORGANIZATION      | ''      | yes     | rmiguelac  | The owner/org of the repo to be Analyzed |
| GH_REPOSITORY        | ''      | yes     | gh-metrics | The repo to be analyzed |
| GH_API_TOKEN         | ''      | yes   | GH-PAT     | The Personal Access Token |

To have them, simply do:

`export VAR=value`

### Run

Simply run it with `go run main.go` or build it with `go build .` then run it `./gh-metrics`


### Output

A `metrics.html` should've been created in the same folder where the code is.

## Example Output

<p align='center'.>
  <img src="./static/metrics.png" width=50% height=50%>
</p>


## TODO

[] Add metrics to influxdb  
[] Add CLI support  
