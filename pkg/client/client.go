package client

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/riversheher/CS4471-trend-service/pkg/models"
)

func GetGainersLosers() (models.GainersLosersResponse, error) {

	var gainersLosersResponse models.GainersLosersResponse

	client := &http.Client{}
	resp, err := client.Get("https://www.alphavantage.co/query?function=TOP_GAINERS_LOSERS&apikey=demo")
	if err != nil {
		return gainersLosersResponse, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return gainersLosersResponse, err
	}

	json.Unmarshal(body, &gainersLosersResponse)

	return gainersLosersResponse, nil
}

func GetGainers(response models.GainersLosersResponse) ([]models.Security, error) {

	var gainers []models.Security
	var timestamp time.Time

	// TODO: Convert the timestamp to a time.Time object
	timestamp = time.Now()

	for _, gainer := range response.Gainers {
		security, err := gainer.ToSecurity(timestamp)
		if err != nil {
			continue
		}
		gainers = append(gainers, *security)
	}

	return gainers, nil
}

func GetLosers(response models.GainersLosersResponse) ([]models.Security, error) {

	var losers []models.Security
	var timestamp time.Time

	// TODO: Convert the timestamp to a time.Time object
	timestamp = time.Now()

	for _, loser := range response.Losers {
		security, err := loser.ToSecurity(timestamp)
		if err != nil {
			continue
		}
		losers = append(losers, *security)
	}

	return losers, nil
}

func GetActive(response models.GainersLosersResponse) ([]models.Security, error) {

	var active []models.Security
	var timestamp time.Time

	// TODO: Convert the timestamp to a time.Time object
	timestamp = time.Now()

	for _, activeSecurity := range response.Active {
		security, err := activeSecurity.ToSecurity(timestamp)
		if err != nil {
			continue
		}
		active = append(active, *security)
	}

	return active, nil
}
