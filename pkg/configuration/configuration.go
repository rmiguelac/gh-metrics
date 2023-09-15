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

	if c.Repository == "" {
		if os.Getenv("GH_REPOSITORY") == "" {
			log.Fatal("Unable to get repository parameter")
		}
		c.Repository = os.Getenv("GH_REPOSITORY")
	}

	log.Println("Reading variables...")
	log.Printf("Organization is %s", c.Organization)
	log.Printf("Repository is %s", c.Repository)

	return &c
}
