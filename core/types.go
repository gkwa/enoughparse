package core

import "fmt"

type GPSInfo struct {
	Latitude  float64
	Longitude float64
}

func (g GPSInfo) String() string {
	return fmt.Sprintf("Latitude: %.6f, Longitude: %.6f", g.Latitude, g.Longitude)
}
