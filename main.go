package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/kakhaUrigashvili/go-rest-api/docs"
	"github.com/kakhaUrigashvili/go-rest-api/handler"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	httpSwagger "github.com/swaggo/http-swagger"
)

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

// Main function
func main() {

	basePath := "/api/v1"

	swaggerHost := getEnv("SWAGGER_HOST", "localhost:8000")

	// programatically set swagger info
	docs.SwaggerInfo.Title = "Rest API Spot Hero"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = swaggerHost
	docs.SwaggerInfo.BasePath = basePath
	docs.SwaggerInfo.Schemes = []string{"http"}

	// Init router
	r := mux.NewRouter()

	// Route handlers
	r.HandleFunc(basePath+"/rates/search", handler.SearchRateHandler).Methods("GET")
	r.HandleFunc(basePath+"/rates", handler.GetRatesHandler).Methods("GET")
	r.HandleFunc(basePath+"/rates", handler.CreateRatesHandler).Methods("POST")
	r.Handle("/metrics", promhttp.Handler()).Methods("GET")
	r.PathPrefix("/").Handler(httpSwagger.Handler(httpSwagger.URL("http://" + docs.SwaggerInfo.Host + "/docs/doc.json")))

	// Start server
	log.Fatal(http.ListenAndServe(":8000", r))
}
