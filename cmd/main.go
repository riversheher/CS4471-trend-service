package main

import (
	"log"
	"net/http"

	"github.com/riversheher/CS4471-trend-service/pkg/handlers"
)

func main() {
	handlers.InitRoutes()

	log.Fatal(http.ListenAndServe(":8080", nil))
}
