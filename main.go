package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kakhaUrigashvili/go-rest-api/handler"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/swaggo/http-swagger"
	"github.com/kakhaUrigashvili/go-rest-api/docs"
)

// Main function
func main() {

	basePath := "/api/v1"

	// programatically set swagger info
	docs.SwaggerInfo.Title = "Rest API Spot Hero"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8000"
	docs.SwaggerInfo.BasePath = basePath
	docs.SwaggerInfo.Schemes = []string{"http"}

	// Init router
	r := mux.NewRouter()

	// Route handlers
	r.HandleFunc(basePath+"/rates/search", handler.SearchRateHandler).Methods("GET")
	r.HandleFunc(basePath+"/rates", handler.GetRatesHandler).Methods("GET")
	r.HandleFunc(basePath+"/rates", handler.CreateRatesHandler).Methods("POST")
	r.Handle("/metrics", promhttp.Handler()).Methods("GET")
	r.PathPrefix("/docs").Handler(httpSwagger.WrapHandler)

	// Start server
	log.Fatal(http.ListenAndServe(":8000", r))
}
