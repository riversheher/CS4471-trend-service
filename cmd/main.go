package main

import (
	"log"
	"net/http"

	"github.com/riversheher/CS4471-trend-service/pkg/web"
)

func main() {
	web.InitRoutes()

	log.Fatal(http.ListenAndServe(":8080", nil))
}
