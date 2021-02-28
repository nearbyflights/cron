package scheduler

import (
	"context"
	"github.com/fagnercarvalho/cron/opensky"
	log "github.com/sirupsen/logrus"
	"sync"
	"time"
)

type Options struct {
	Interval time.Duration
}

func GetFlights(ctx context.Context, wg *sync.WaitGroup, options Options) (<-chan []opensky.Flight, error) {
	flightsCh := make(chan []opensky.Flight)
	ticker := time.NewTicker(options.Interval)

	go func() {
		defer wg.Done()

		for {
			select {
			case <-ticker.C:
				flights, err := getFlights(ctx)
				if err != nil {
					log.Error(err)
					continue
				}

				flightsCh <- flights
			case <-ctx.Done():
				log.Info("Program stopped: finish get flights routine")
				return
			}
		}
	}()

	return flightsCh, nil
}

func getFlights(ctx context.Context) ([]opensky.Flight, error) {
	flights, err := opensky.GetFlights(ctx)
	if err != nil {
		return nil, err
	}

	log.Infof("OpenSky API response length: %v bytes", len(flights))

	parsedFlights, err := opensky.ParseFlights(flights)
	if err != nil {
		return nil, err
	}

	return parsedFlights, nil
}
