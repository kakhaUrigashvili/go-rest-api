package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kakhaUrigashvili/go-rest-api/handler"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Main function
func main() {
	// Init router
	r := mux.NewRouter()

	routePrefix := "/api/v1"

	// Route handlers
	r.HandleFunc(routePrefix+"/rates/search", handler.SearchRateHandler).Methods("GET")
	r.HandleFunc(routePrefix+"/rates", handler.GetRatesHandler).Methods("GET")
	r.HandleFunc(routePrefix+"/rates", handler.CreateRatesHandler).Methods("POST")
	r.Handle("/metrics", promhttp.Handler()).Methods("GET")

	// Start server
	log.Fatal(http.ListenAndServe(":8000", r))
}
