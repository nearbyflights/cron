package opensky

import (
	"context"
	"io/ioutil"
	"net/http"
)

var tr = &http.Transport{
	MaxConnsPerHost: 1,
}

var client = &http.Client{
	Transport: tr,
}

func GetFlights(ctx context.Context) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", "https://opensky-network.org/api/states/all", nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
