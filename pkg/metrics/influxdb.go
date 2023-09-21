package metrics

import (
	"log"

	influxdb2 "github.com/influxdata/influxdb-client-go"
)

type InfluxConfig struct {
	Bucket       string
	Host         string
	Port         string
	Auth         string
	Organization string
}

type InfluxStore struct {
	client *influxdb2.Client
}

func New(c *InfluxConfig) (*InfluxStore, error) {

	log.Printf("Connecting to bucket %s, in %s", c.Bucket, c.Host)
	client := influxdb2.NewClient("http://"+c.Host+":"+c.Port, c.Auth)

	return &InfluxStore{
		client: &client,
	}, nil

}
