package configuration

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

type Configuration struct {
	Organization string
	Repository   string
	Report       *Report
	TimeseriesDB *TimeseriesDB
}

type TimeseriesDB struct {
	Host         string
	Port         string
	Auth         string
	Bucket       string
	Organization string
}

type Report struct {
	Colors *Colors
	Data   *Data
}

type Colors struct {
	Failure string
	Success string
}

type Data struct {
	Last *Last
}

type Last struct {
	Days   int
	Hours  int
	Months int
	Weeks  int
}

func New() *Configuration {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetConfigType("yaml")

	viper.SetDefault("report.data.last.days", 7)
	viper.SetDefault("report.colors.success", "139, 193, 71, 0.7")
	viper.SetDefault("report.colors.failure", "255, 61, 71, 0.7")
	viper.SetDefault("timeseriesdb.auth", os.Getenv("GH_INFLUX_TOKEN"))

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	var c Configuration
	err := viper.Unmarshal(&c)
	if err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}

	if c.Organization == "" {
		if os.Getenv("GH_ORGANIZATION") == "" {
			log.Fatal("Unable to get oganization parameter")
		}
		c.Organization = os.Getenv("GH_ORGANIZATION")
	}

	if os.Getenv("GH_ORGANIZATION") != "" {
		c.Organization = os.Getenv("GH_ORGANIZATION")
	}

	if c.Repository == "" {
		if os.Getenv("GH_REPOSITORY") == "" {
			log.Fatal("Unable to get repository parameter")
		}
		c.Repository = os.Getenv("GH_REPOSITORY")
	}

	if os.Getenv("GH_REPOSITORY") != "" {
		c.Organization = os.Getenv("GH_REPOSITORY")
	}

	log.Println("Reading configuration file...")
	log.Printf("Organization is %s", c.Organization)
	log.Printf("Repository is %s", c.Repository)

	return &c
}
