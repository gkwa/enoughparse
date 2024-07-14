package core

import (
	"encoding/json"
	"fmt"

	"github.com/go-logr/logr"
)

func ProcessImage(imagePath, outputFormat string, logger logr.Logger) error {
	gpsInfo, err := ExtractGPSCoordinates(imagePath, logger)
	if err != nil {
		return fmt.Errorf("failed to extract GPS coordinates: %w", err)
	}

	logger.V(1).Info("Extracted GPS coordinates", "gpsInfo", gpsInfo.String())

	link := GenerateGoogleMapsLink(gpsInfo)
	coordinates := fmt.Sprintf("latitude:%.6f longitude:%.6f", gpsInfo.Latitude, gpsInfo.Longitude)

	switch outputFormat {
	case "json":
		output := map[string]interface{}{
			"link": link,
			"coordinates": map[string]float64{
				"latitude":  gpsInfo.Latitude,
				"longitude": gpsInfo.Longitude,
			},
			"image": imagePath,
		}
		jsonOutput, err := json.Marshal(output)
		if err != nil {
			return fmt.Errorf("failed to marshal JSON: %w", err)
		}
		fmt.Println(string(jsonOutput))
	case "text":
		fmt.Printf("%s %s %s\n", link, coordinates, imagePath)
	default:
		return fmt.Errorf("unsupported output format: %s", outputFormat)
	}

	return nil
}
