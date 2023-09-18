package metrics

import (
	"context"
	"log"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go"
	"github.com/rmiguelac/gh-metrics/pkg/configuration"
)

func ToInfluxDB(c *configuration.Configuration, m *Metrics) error {

	log.Printf("Connecting to bucket %s, in %s", c.TimeseriesDB.Bucket, c.TimeseriesDB.Host)
	client := influxdb2.NewClient("http://"+c.TimeseriesDB.Host+":"+c.TimeseriesDB.Port, c.TimeseriesDB.Auth)
	writeAPI := client.WriteAPIBlocking(c.TimeseriesDB.Organization, c.TimeseriesDB.Bucket)
	p := influxdb2.NewPoint("stat",
		map[string]string{"unit": "temperature"},
		map[string]interface{}{"avg": 24.5, "max": 45},
		time.Now())
	// write point immediately
	err := writeAPI.WritePoint(context.Background(), p)
	if err != nil {
		log.Printf("Unable to push to influxdb. %s", err)
		return err
	}
	// create point using fluent style
	p = influxdb2.NewPointWithMeasurement("stat").
		AddTag("unit", "temperature").
		AddField("avg", 23.2).
		AddField("max", 45).
		SetTime(time.Now())
	writeAPI.WritePoint(context.Background(), p)
	client.Close()

	return nil
}
