package opensky

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Flight struct {
	Latitude  float64
	Longitude float64
	Country   string
	CallSign  string
	Icao24    string
	Velocity  float64
}

type openSkyFlight struct {
	Time   int             `json:"time"`
	States [][]interface{} `json:"states"`
}

func ParseFlights(jsonBytes []byte) ([]Flight, error) {
	var openSkyFlights openSkyFlight

	err := json.Unmarshal(jsonBytes, &openSkyFlights)
	if err != nil {
		return []Flight{}, err
	}

	var flights []Flight
	for _, f := range openSkyFlights.States {
		callSign := strings.TrimSpace(fmt.Sprintf("%v", f[1]))

		if callSign == "" {
			continue
		}

		flight := Flight{
			Latitude:  parseFloat(f[6]),
			Longitude: parseFloat(f[5]),
			Country:   fmt.Sprintf("%v", f[2]),
			CallSign:  callSign,
			Icao24:    fmt.Sprintf("%v", f[0]),
			Velocity:  parseFloat(f[9]),
		}

		flights = append(flights, flight)
	}

	return flights, nil
}

func parseFloat(value interface{}) float64 {
	if v, ok := value.(float64); ok {
		return v
	}

	return 0
}
