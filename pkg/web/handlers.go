package web

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/riversheher/CS4471-trend-service/pkg/client"
	"github.com/riversheher/CS4471-trend-service/pkg/registration"
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

	json, err := json.Marshal(gainers)
	if err != nil {
		http.Error(w, "Failed to convert", http.StatusInternalServerError)
		return
	}
	w.Write(json)
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

	json, err := json.Marshal(losers)
	if err != nil {
		http.Error(w, "Failed to convert", http.StatusInternalServerError)
		return
	}

	w.Write(json)
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

	json, err := json.Marshal(active)
	if err != nil {
		http.Error(w, "Failed to convert", http.StatusInternalServerError)
		return
	}

	w.Write(json)
}

func (app *Application) Register(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Loading configuration...")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting to port %s", port)
	}

	confFile, err := os.Open("config.json")
	if err != nil {
		log.Fatal(err)
	}
	defer confFile.Close()

	conf, err := io.ReadAll(confFile)
	if err != nil {
		log.Fatal(err)
	}

	var config map[string]string
	err = json.Unmarshal(conf, &config)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Configuration Loaded...")

	fmt.Println("Initiating registration...")
	// Get the token from the registry
	token, err := registration.GetTokenFromRegistry(config["registryURL"])
	if err != nil {
		fmt.Println("Failed to get token from registry")
		fmt.Println(err)
	}
	// Register the service with the registry
	appInfo := map[string]string{
		"serviceName": config["serviceName"],
		"port":        port,
		"description": config["description"],
		"version":     config["version"],
		"instanceId":  config["instanceId"],
		"url":         config["url"],
	}
	response, err := registration.RegisterSelf(config["registryURL"], token.(string), appInfo)
	if err != nil {
		fmt.Println("Failed to register service with registry")
		fmt.Println(err)
	}

	fmt.Println(response)

	fmt.Println("Registration complete...")
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
	http.HandleFunc("/register", app.Register)
	http.HandleFunc("/", app.Home)
}
