package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

// build crontab job accordingly to the sunset

type Envelop struct {
	Results Results `json:"results"`
	Status  string  `json:"status"`
}

type Results struct {
	Sunrise   string `json:"sunrise"`
	Sunset    string `json:"sunset"`
	DayLength int    `json:"day_length"`
}

func main() {

	stdout := log.New(os.Stdout, "", log.Ltime)
	stdout.Println("Creating cron job starting from today up to 1 year")

	i := 1
	nDays := 17
	d := time.Now()
	tzRome, _ := time.LoadLocation("Europe/Rome")

	debug := false
	if len(os.Args) == 2 && os.Args[1] == "d" {
		stdout.Println("*** DEBUG ***")
		debug = true
	}

	f, e := os.Create("datcrony")
	if e != nil {
		stdout.Fatal(e)
	}

	defer f.Close()

	for {
		i += 1
		if i == nDays {
			stdout.Println("done")
			break
		}

		d = d.Add(24 * time.Hour)

		// Get the sunset time
		res, err := http.Get(fmt.Sprintf(
			"https://api.sunrise-sunset.org/json?lng=12.501337544365304&lat=41.96905141478927&formatted=0&date=%s",
			d.Format("2006-01-02")))
		if err != nil {
			stdout.Fatal(err)
		}

		if res.StatusCode != http.StatusOK {
			stdout.Fatalf("got HTTP %v for %s ", res.StatusCode, d.Format("2006-01-02"))
		}

		var envelop Envelop
		data, _ := ioutil.ReadAll(res.Body)
		json.Unmarshal(data, &envelop)

		sunset, err := time.Parse("2006-01-02T15:04:05-07:00", envelop.Results.Sunset)
		if err != nil {
			stdout.Fatal(err)
		}

		sunset = sunset.In(tzRome)

		if debug {
			sunset = time.Now().Add(2 * time.Minute)
		}

		cronString := strings.Join([]string{
			sunset.Format("04"), // minute
			sunset.Format("15"), // hour
			sunset.Format("02"), // day
			sunset.Format("01"), // month
			"*",
			"/home/pi/sunsetty/run.sh",
		}, " ")

		stdout.Printf("Day : %s", sunset.String())
		stdout.Printf("Cron: %s", cronString)

		f.Write([]byte(cronString + "\n"))

		if debug {
			stdout.Println("debug - bye")
			break
		}
	}

	f.Write([]byte("\n"))
	stdout.Println("Rember to update crontab!")
	stdout.Println("crontab -r && crontab datcrony")
}
