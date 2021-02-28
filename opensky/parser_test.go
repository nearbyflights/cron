package opensky

import (
	"testing"
)

func TestParseFlights(t *testing.T) {
	flights, err := ParseFlights([]byte("{\"time\":1608491050,\"states\":[[\"e47f51\",\"GLO1335 \",\"Brazil\",1608491044,1608491050,-46.6582,-23.6243,null,true,97.31,326.65,-5.53,null,null,null,false,0]]}"))
	if err != nil {
		t.Fatal("error while parsing flights JSON")
	}

	flight := flights[0]

	if flight.Longitude != -46.6582 {
		t.Errorf("excepted %v but received %v in longitude field", -46.6582, flight.Longitude)
	}

	if flight.Latitude != -23.6243 {
		t.Errorf("excepted %v but received %v in latitude field", -23.6243, flight.Latitude)
	}

	if flight.Velocity != 97.31 {
		t.Errorf("excepted %v but received %v in velocity field", 97.31, flight.Velocity)
	}

	if flight.Country != "Brazil" {
		t.Errorf("excepted %v but received %v in country field", "Brazil", flight.Country)
	}

	if flight.Icao24 != "e47f51" {
		t.Errorf("excepted %v but received %v in ICAO 24 field", "e47f51", flight.Icao24)
	}

	if flight.CallSign != "GLO1335" {
		t.Errorf("excepted %v but received %v in call sign field", "GLO1335", flight.CallSign)
	}
}

func TestGetFlights_Multiple(t *testing.T) {
	flights, err := ParseFlights([]byte("{\"time\":1608494070,\"states\":[[\"e490d3\",\"AZU4004 \",\"Brazil\",1608494061,1608494069,-46.6607,-23.6205,null,true,2.83,146.25,null,null,null,null,false,0],[\"e4819c\",\"TAM3143 \",\"Brazil\",1608493959,1608493959,-46.6575,-23.6251,792.48,false,76.91,146.74,0.65,null,800.1,null,false,0]]}"))
	if err != nil {
		t.Fatal("error while parsing flights JSON")
	}

	if 2 != len(flights) {
		t.Fatalf("expected length %v but received %v", 2, len(flights))
	}
}

func TestGetFlights_NoResults(t *testing.T) {
	flights, err := ParseFlights([]byte("{\"time\":1608495350,\"states\":null}"))
	if err != nil {
		t.Fatal("error while parsing flights JSON")
	}

	if 0 != len(flights) {
		t.Fatalf("expected length %v but received %v", 0, len(flights))
	}
}

func TestGetFlights_EmptyJson(t *testing.T) {
	flights, err := ParseFlights([]byte("{}"))
	if err != nil {
		t.Fatal("error while parsing flights JSON")
	}

	if 0 != len(flights) {
		t.Fatalf("expected length %v but received %v", 0, len(flights))
	}
}

func TestGetFlights_Invalid(t *testing.T) {
	_, err := ParseFlights([]byte("{"))
	if err == nil {
		t.Fatal("expected error while parsing flights JSON")
	}
}
