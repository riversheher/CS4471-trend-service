package web

import (
	"fmt"
	"net/http"

	"github.com/riversheher/CS4471-trend-service/pkg/client"
)

func (app *Application) Home(w http.ResponseWriter, r *http.Request) {
	app.Render(w, "home.html")
}

func (app *Application) Gainers(w http.ResponseWriter, r *http.Request) {
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

func (app *Application) Losers(w http.ResponseWriter, r *http.Request) {
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

func (app *Application) MostActive(w http.ResponseWriter, r *http.Request) {
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
	app := NewApplication()
	app.Initialize()
	if err := app.Initialize(); err != nil {
		panic(err)
	}

	http.HandleFunc("/gainers", app.Gainers)
	http.HandleFunc("/losers", app.Losers)
	http.HandleFunc("/active", app.MostActive)
	http.HandleFunc("/", app.Home)
}
