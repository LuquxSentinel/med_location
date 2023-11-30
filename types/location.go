package types

type Location struct {
	Type        string    `json:"Point"`
	Coordinates []float64 `json:"coordinates"`
}
