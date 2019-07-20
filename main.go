package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kakhaUrigashvili/go-rest-api/handler"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/swaggo/http-swagger"
	_"github.com/kakhaUrigashvili/go-rest-api/docs"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server celler server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

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
	r.PathPrefix("/docs").Handler(httpSwagger.WrapHandler)

	// Start server
	log.Fatal(http.ListenAndServe(":8000", r))
}
