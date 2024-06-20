package handlers

import (
	"encoding/json"
	"net/http"
	"spaceVoyagers/models"
	"spaceVoyagers/store"
	"spaceVoyagers/utils"
	"strconv"

	"github.com/gorilla/mux"
)

var Store = store.NewExoplanet()

func AddExoplanetHandlerFunc(w http.ResponseWriter, r *http.Request) {

	exoplanetModel := models.ExoplanetModel{}
	err := json.NewDecoder(r.Body).Decode(&exoplanetModel)
	if err != nil {
		http.Error(w, "Error in decoding the request data", http.StatusBadGateway)
		return
	}

	exoplanetModel.ID = utils.UUID()
	Store.AddExoplanet(exoplanetModel)
	w.Write([]byte("Exoplanet Added Sucessfully"))
	w.WriteHeader(http.StatusCreated)

}

func ListExoplanetHandlerFunc(w http.ResponseWriter, r *http.Request) {
	explonet := Store.ListExoplanet()
	w.Write([]byte("Exoplanet Listed Sucessfully"))
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(explonet)
}

func DeleteExoplanetHandlerFunc(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	id := v["id"]
	err := Store.DeleteExoplanet(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}
	w.Write([]byte("Exoplanet Deleted Sucessfully"))
	w.WriteHeader(http.StatusNoContent)
}

func UpdateExoplanetHandlerFunc(w http.ResponseWriter, r *http.Request) {
	exoplanetModel := models.ExoplanetModel{}
	if err := json.NewDecoder(r.Body).Decode(&exoplanetModel); err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
	}

	v := mux.Vars(r)
	key := v["id"]
	if err := Store.UpdateExoplanet(key, exoplanetModel); err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}

	w.Write([]byte("Exoplanet Updated Sucessfully"))
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(exoplanetModel)

}

func GetExoplanetByIdHandlerFunc(w http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)
	id := pathParams["id"]
	exoplanet, err := Store.GetExoplanetById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}
	w.Write([]byte("Exoplanet get Sucessfully"))
	w.WriteHeader(http.StatusNoContent)
	json.NewEncoder(w).Encode(exoplanet)
}

func FuelEstimationHandlerFunc(w http.ResponseWriter, r *http.Request) {
	// get the explonet type based on id,
	pathParams := mux.Vars(r)
	id := pathParams["id"]

	crewCapacity, err := strconv.Atoi(r.URL.Query().Get("crewCapacity"))
	if err != nil || crewCapacity <= 0 {
		http.Error(w, "invalid crew capacity", http.StatusBadRequest)
		return
	}

	exoplanet, err := Store.GetExoplanetById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}

	// find gravity for that type
	// find fuel estimation
	fuel, err := utils.FuelEstimation(exoplanet, crewCapacity)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]float64{"fuel_estimation In Units": fuel})

}
