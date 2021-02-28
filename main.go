package main

import (
	"context"
	"github.com/fagnercarvalho/cron/db"
	"github.com/fagnercarvalho/cron/scheduler"
	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type Configuration struct {
	PostgresUrl  string `required:"true" envconfig:"POSTGRES_URL" default:"localhost:5432"`
	User         string `required:"true" envconfig:"POSTGRES_USER" default:"admin"`
	Password     string `required:"true" envconfig:"POSTGRES_PASSWORD" default:"secret"`
	DatabaseName string `required:"true" envconfig:"POSTGRES_DB" default:"flights"`
}

func init() {
	log.SetOutput(os.Stdout)
	log.SetFormatter(&log.TextFormatter{ForceColors: true, FullTimestamp: true})
}

func main() {
	var c Configuration
	err := envconfig.Process("", &c)
	if err != nil {
		log.Fatal(err.Error())
	}

	database := db.ClientOptions{
		Address:  c.PostgresUrl,
		User:     c.User,
		Password: c.Password,
		Database: c.DatabaseName,
	}

	client := db.NewClient(database)

	signals := make(chan os.Signal)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(context.Background())

	wg := sync.WaitGroup{}

	options := scheduler.Options{
		Interval: time.Second * 15,
	}

	wg.Add(1)
	flights, err := scheduler.GetFlights(ctx, &wg, options)
	if err != nil {
		log.Error(err)
	}

	go func() {
		for {
			select {
			case flights := <-flights:
				log.Infof("Flight count: %v", len(flights))

				var newFlights []db.NewFlight
				for _, flight := range flights {
					newFlights = append(newFlights, db.NewFlight{
						Latitude:  flight.Latitude,
						Longitude: flight.Longitude,
						Country:   flight.Country,
						CallSign:  flight.CallSign,
						Icao24:    flight.Icao24,
						Velocity:  flight.Velocity,
					})
				}

				log.Info("Saving flights on database")
				err := client.SaveFlights(newFlights)
				if err != nil {
					log.Error(err)
				}
			}
		}
	}()

	<-signals
	log.Info("Closing db connection")
	client.Close()
	cancel()
	log.Info("Cancelling context")
	wg.Wait()
	log.Info("Finishing")
	os.Exit(0)
}
