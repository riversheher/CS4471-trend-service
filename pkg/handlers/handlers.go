package handlers

import (
	"fmt"
	"net/http"

	"github.com/riversheher/CS4471-trend-service/pkg/client"
)

func Gainers(w http.ResponseWriter, r *http.Request) {
	response, err := client.GetGainersLosers()
	if err != nil {
		http.Error(w, "Failed to fetch data", http.StatusInternalServerError)
		return
	}
	gainers, err := client.GetGainers(response)
	if err != nil {
		http.Error(w, "Failed to convert", http.StatusInternalServerError)
		return
	}
	fmt.Fprint(w, gainers)
}

func Losers(w http.ResponseWriter, r *http.Request) {
	response, err := client.GetGainersLosers()
	if err != nil {
		http.Error(w, "Failed to fetch data", http.StatusInternalServerError)
		return
	}
	losers, err := client.GetLosers(response)
	if err != nil {
		http.Error(w, "Failed to convert", http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, losers)
}

func MostActive(w http.ResponseWriter, r *http.Request) {
	response, err := client.GetGainersLosers()
	if err != nil {
		http.Error(w, "Failed to fetch data", http.StatusInternalServerError)
		return
	}
	active, err := client.GetActive(response)
	if err != nil {
		http.Error(w, "Failed to convert", http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, active)
}

func InitRoutes() {
	http.HandleFunc("/gainers", Gainers)
	http.HandleFunc("/losers", Losers)
	http.HandleFunc("/active", MostActive)
}
