package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"spaceVoyagers/models"
	"testing"

	"github.com/gorilla/mux"
)

// func setupRouter() *mux.Router {
// 	r := mux.NewRouter()
// 	r.HandleFunc("/fuelEstimation/{id}", .FuelEstimationHandlerFunc).Methods("GET")
// 	return r
// }

func TestCreateExoplanet(t *testing.T) {
	// Prepare test data
	exoplanet := models.ExoplanetModel{
		Name:        "Test Planet",
		Description: "A test planet",
		Distance:    50,
		Radius:      1.5,
		Type:        "GasGiant",
	}
	body, _ := json.Marshal(exoplanet)

	// Create a new HTTP request
	req, err := http.NewRequest("POST", "/addExoplanet", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Create a new handler and call it
	handler := http.HandlerFunc(AddExoplanetHandlerFunc)
	handler.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}

	// Check the response body
	var responseExoplanet models.ExoplanetModel
	json.Unmarshal(rr.Body.Bytes(), &responseExoplanet)
	if responseExoplanet.Name != exoplanet.Name {
		t.Errorf("handler returned unexpected body: Receive %v want %v",
			responseExoplanet.Name, exoplanet.Name)
	}
}

func TestListExoplanets(t *testing.T) {
	req, err := http.NewRequest("GET", "/listExoplanet", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ListExoplanetHandlerFunc)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestGetExoplanet(t *testing.T) {
	exoplanet := models.ExoplanetModel{
		Name:        "Test Planet",
		Description: "A test planet",
		Distance:    50,
		Radius:      1.5,
		Mass:        2.0,
		Type:        models.GasGiant,
	}
	Store.AddExoplanet(exoplanet)

	req, err := http.NewRequest("GET", "/getExoplanet/"+exoplanet.ID, nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	r := mux.NewRouter()
	r.HandleFunc("/getExoplanet/{id}", GetExoplanetByIdHandlerFunc)
	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var responseExoplanet models.ExoplanetModel
	json.Unmarshal(rr.Body.Bytes(), &responseExoplanet)
	if responseExoplanet.Name != exoplanet.Name {
		t.Errorf("handler returned unexpected body: got %v want %v",
			responseExoplanet.Name, exoplanet.Name)
	}
}

func TestUpdateExoplanet(t *testing.T) {
	exoplanet := models.ExoplanetModel{
		Name:        "Test Planet",
		Description: "A test planet",
		Distance:    50,
		Radius:      1.5,
		Type:        models.GasGiant,
	}
	Store.AddExoplanet(exoplanet)

	updatedExoplanet := exoplanet
	updatedExoplanet.Description = "An updated test planet"
	body, _ := json.Marshal(updatedExoplanet)

	req, err := http.NewRequest("PUT", "/exoplanets/"+exoplanet.ID, bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	r := mux.NewRouter()
	r.HandleFunc("/updateExoplanet/{id}", UpdateExoplanetHandlerFunc)
	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var responseExoplanet models.ExoplanetModel
	json.Unmarshal(rr.Body.Bytes(), &responseExoplanet)
	if responseExoplanet.Description != updatedExoplanet.Description {
		t.Errorf("handler returned unexpected body: got %v want %v",
			responseExoplanet.Description, updatedExoplanet.Description)
	}
}

func TestDeleteExoplanet(t *testing.T) {
	exoplanet := models.ExoplanetModel{
		Name:        "Test Planet",
		Description: "A test planet",
		Distance:    50,
		Radius:      1.5,
		Type:        models.GasGiant,
	}
	Store.AddExoplanet(exoplanet)

	req, err := http.NewRequest("DELETE", "/deleteExoplanet/"+exoplanet.ID, nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	r := mux.NewRouter()
	r.HandleFunc("/deleteExoplanet/{id}", DeleteExoplanetHandlerFunc)
	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNoContent {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNoContent)
	}
}

func TestEstimateFuel(t *testing.T) {
	exoplanet := models.ExoplanetModel{
		Name:        "Test Planet",
		Description: "A test planet",
		Distance:    50,
		Radius:      1.5,
		Type:        models.GasGiant,
	}
	Store.AddExoplanet(exoplanet)

	// Setup the router and create a test server
	// router := setupRouter()
	// server := httptest.NewServer(router)
	// defer server.Close()

	req, err := http.NewRequest("GET", "/fuelEstimation/"+exoplanet.ID+"?crewCapacity=10", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	r := mux.NewRouter()
	r.HandleFunc("/fuelEstimation/{id}", FuelEstimationHandlerFunc)
	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var response map[string]float64
	json.Unmarshal(rr.Body.Bytes(), &response)
	expectedFuel := float64(50) / (0.5 / 1.5 * 1.5) * 10
	if response["fuel_estimation"] != expectedFuel {
		t.Errorf("handler returned unexpected body: got %v want %v",
			response["fuel_estimation"], expectedFuel)
	}
}
