package db

import (
	"fmt"
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/types"
	log "github.com/sirupsen/logrus"
)

type Flight struct {
	Id        int     `sql:"id"`
	Geometry  types.Q `sql:"geom"`
	Latitude  float64 `sql:"latitude"`
	Longitude float64 `sql:"longitude"`
	Country   string  `sql:"country"`
	CallSign  string  `sql:"call_sign"`
	Icao24    string  `sql:"icao"`
	Velocity  float64 `sql:velocity`
}

type NewFlight struct {
	Geometry  string
	Latitude  float64
	Longitude float64
	Country   string
	CallSign  string
	Icao24    string
	Velocity  float64
}

type Client struct {
	database *pg.DB
}

type ClientOptions struct {
	Address  string
	User     string
	Password string
	Database string
}

func NewClient(options ClientOptions) Client {
	db := pg.Connect(&pg.Options{
		Addr:     options.Address,
		User:     options.User,
		Password: options.Password,
		Database: options.Database,
	})

	return Client{db}
}

func (c *Client) SaveFlights(newFlights []NewFlight) error {
	var flights []Flight
	for _, newFlight := range newFlights {
		flights = append(flights, Flight{
			Geometry:  types.Q(fmt.Sprintf("ST_SetSRID(ST_MakePoint(%v, %v),4326)", newFlight.Longitude, newFlight.Latitude)),
			Latitude:  newFlight.Latitude,
			Longitude: newFlight.Longitude,
			Country:   newFlight.Country,
			CallSign:  newFlight.CallSign,
			Icao24:    newFlight.Icao24,
			Velocity:  newFlight.Velocity,
		})
	}

	var existingFlights []Flight
	err := c.database.Model(&existingFlights).Select()
	if err != nil {
		return err
	}

	_, err = c.database.Exec("TRUNCATE flights CASCADE;")
	if err != nil {
		return err
	}

	_, err = c.database.Model(&flights).Insert()
	if err != nil {
		return err
	}

	log.Infof("Inserted %v new flights", len(flights))

	return nil
}

func (c *Client) Close() {
	c.database.Close()
}