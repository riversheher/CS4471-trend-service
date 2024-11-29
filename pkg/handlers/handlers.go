package handlers

import (
	"fmt"
	"net/http"
)

func Gainers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Top Gainers")
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
