package main

import (
	"log"
	"net/http"

	"spaceVoyagers/handlers"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/addExoplanet", handlers.AddExoplanetHandlerFunc).Methods("POST")
	r.HandleFunc("/listExoplanet", handlers.ListExoplanetHandlerFunc).Methods("GET")
	r.HandleFunc("/getExoplanet/{id}", handlers.GetExoplanetByIdHandlerFunc).Methods("GET")
	r.HandleFunc("/updateExoplanet/{id}", handlers.UpdateExoplanetHandlerFunc).Methods("PUT")
	r.HandleFunc("/deleteExoplanet/{id}", handlers.DeleteExoplanetHandlerFunc).Methods("DELETE")
	r.HandleFunc("/fuelEstimation/{id}", handlers.FuelEstimationHandlerFunc).Methods("GET")

	log.Println("Starting server on :8080")
	http.ListenAndServe(":8080", r)
}
