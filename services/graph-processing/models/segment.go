package models

type Segment struct {
	ID        string  // Segment ID
	Source    string  // Source node ID
	Target    string  // Target node ID
	Distance  float64 // Distance or weight
}