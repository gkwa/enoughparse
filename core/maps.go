package core

import (
	"fmt"
)

func GenerateGoogleMapsLink(gpsInfo *GPSInfo) string {
	return fmt.Sprintf("https://www.google.com/maps?q=%.6f,%.6f", gpsInfo.Latitude, gpsInfo.Longitude)
}
