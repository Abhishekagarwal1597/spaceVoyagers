package models

type ExoplanetModel struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Distance    float64 `json:"distance"`
	Radius      float64 `json:"radius"`
	Mass        float64 `json:"mass"`
	Type        string  `json:"type"`
}

// ExoplanetType is a type for defining the type of exoplanet.
// type ExoplanetType string

const (
	GasGiant    string = "GasGiant"
	Terrestrial string = "Terrestrial"
)
