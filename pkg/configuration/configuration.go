package configuration

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

type Configuration struct {
	Organization string
	Repository   string
}

func New() *Configuration {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	var c Configuration
	viper.SetDefault(c.Organization, os.Getenv("GH_ORGANIZATION"))
	viper.SetDefault(c.Repository, os.Getenv("GH_REPOSITORY"))
	err := viper.Unmarshal(&c)
	if err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}

	log.Fatalf("Reading variables...")
	log.Fatalf("Organization is %s", c.Organization)
	log.Fatalf("Repository is %s", c.Repository)

	return &c
}
