package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kakhaUrigashvili/go-rest-api/handler"
)

// Main function
func main() {
	// Init router
	r := mux.NewRouter()

	routePrefix := "/api/v1"

	// Route handles & endpoints
	r.HandleFunc(routePrefix+"/rates/search", handler.RateSearch).Methods("GET")
	r.HandleFunc(routePrefix+"/rates", handler.GetRates).Methods("GET")
	r.HandleFunc(routePrefix+"/rates", handler.CreateRates).Methods("POST")

	// Start server
	log.Fatal(http.ListenAndServe(":8000", r))
}
