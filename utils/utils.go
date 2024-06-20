package utils

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math"
	"spaceVoyagers/models"
)

func UUID() string {
	b := make([]byte, 16)
	_, _ = rand.Read(b)
	appId := fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	//Application Id generated
	return appId
}

// FuelEstimation estimates the fuel required to reach an exoplanet.
func FuelEstimation(exoplanet models.ExoplanetModel, crewCapacity int) (float64, error) {
	var gravity float64
	if exoplanet.Type == "GasGiant" {
		gravity = 0.5 / math.Pow(exoplanet.Radius, 2)
	} else if exoplanet.Type == "Terrestrial" && exoplanet.Mass == 0 {
		gravity = exoplanet.Mass / math.Pow(exoplanet.Radius, 2)
	} else {
		return 0, errors.New("invalid exoplanet data for gravity calculation")
	}
	return float64(exoplanet.Distance) / math.Pow(gravity, 2) * float64(crewCapacity), nil
}
