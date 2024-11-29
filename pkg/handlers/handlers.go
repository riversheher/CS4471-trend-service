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
	fmt.Fprint(w, response)
}

func Losers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Top Losers")
}

func MostActive(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Most Active")
}

func InitRoutes() {
	http.HandleFunc("/gainers", Gainers)
	http.HandleFunc("/losers", Losers)
	http.HandleFunc("/most-active", MostActive)
}
