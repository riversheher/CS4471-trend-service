package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/riversheher/CS4471-trend-service/pkg/registration"
	"github.com/riversheher/CS4471-trend-service/pkg/web"
)

func main() {
	web.InitRoutes()

	fmt.Println("Loading configuration...")

	confFile, err := os.Open("config.json")
	if err != nil {
		log.Fatal(err)
	}
	defer confFile.Close()

	conf, err := io.ReadAll(confFile)
	if err != nil {
		log.Fatal(err)
	}

	var config map[string]interface{}
	err = json.Unmarshal(conf, &config)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Configuration Loaded...")

	fmt.Println("Initiating registration...")
	// Get the token from the registry
	token, err := registration.GetTokenFromRegistry(config["registryURL"].(string))
	if err != nil {
		fmt.Println("Failed to get token from registry")
	}
	// Register the service with the registry
	appInfo := map[string]interface{}{
		"serviceName": config["serviceName"],
		"port":        config["port"],
		"description": config["description"],
		"version":     config["version"],
		"instanceId":  config["instanceId"],
		"url":         config["url"],
	}
	_, err = registration.RegisterSelf(config["registryURL"].(string), token.(string), appInfo)
	if err != nil {
		fmt.Println("Failed to register service with registry")
	}
	fmt.Println("Registration complete...")

	fmt.Println("Starting server on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
