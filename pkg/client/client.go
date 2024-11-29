package client

import (
	"encoding/json"
	"io"
	"net/http"

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
