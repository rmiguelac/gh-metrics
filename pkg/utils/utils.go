package utils

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/rmiguelac/gh-metrics/pkg/configuration"
	"github.com/spf13/viper"
)

/*
	GetLastNDateTimeBack accepts a string for 'weeks', 'days' or 'hours'

and an int of how many of those date types to go back
*/
func getLastNDateTimeBack(s string, d int) (string, error) {

	t := time.Now()
	var delta int

	/* add format expected by gh */
	switch s {
	case "months":
		return t.AddDate(0, -d, 0).Format(time.RFC3339), nil
	case "weeks":
		return t.AddDate(0, 0, -d*7).Format(time.RFC3339), nil
	case "days":
		delta = d * 24 * int(time.Hour)
		return t.AddDate(0, 0, -d).Format(time.RFC3339), nil
	case "hours":
		dur, _ := time.ParseDuration(strconv.Itoa(d) + "h")
		return t.Add(-dur).Format(time.RFC3339), nil
	}

	return t.AddDate(0, 0, -delta).String(), errors.New("Unable to get past date")
}

func getToday() string {
	return time.Now().Format(time.RFC3339)
}

func GetTimeRange(c *configuration.Configuration) (string, error) {
	t, a, err := getLastDateTimeAmount(c)
	if err != nil {
		log.Printf("Unable to get how much time back")
		return "", err
	}
	dateback, err := getLastNDateTimeBack(t, a)
	if err != nil {
		log.Printf("Unable to get precise date from %d %s back.", a, t)
		return "", err
	}
	tdy := getToday()
	trange := fmt.Sprintf("%s..%s", dateback, tdy)

	log.Printf("Time range is between %s and %s", dateback, tdy)
	return trange, nil
}

/* Read configuration and reture the key under report.data.last */
func getLastDateTimeAmount(c *configuration.Configuration) (string, int, error) {

	ldta := viper.GetStringMap("report.data.last")
	for k, v := range ldta {
		amount, _ := strconv.Atoi(fmt.Sprintf("%v", v))
		return k, amount, nil
	}

	return "weeks", 1, errors.New("Unable to get last date time amount. Defaulting to a week.")
}
