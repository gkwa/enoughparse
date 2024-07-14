package core

import (
	"fmt"

	"github.com/go-logr/logr"
)

func ExtractGPSCoordinates(imagePath string, logger logr.Logger) (*GPSInfo, error) {
	logger.V(2).Info("Parsing GPS coordinates")

	latStr, lonStr, lonRef, err := extractGPSInfo(imagePath, logger)
	if err != nil {
		return nil, err
	}

	logger.V(2).Info("Parsing latitude", "latStr", latStr)
	lat, err := parseCoordinate(latStr, logger)
	if err != nil {
		logger.Error(err, "Failed to parse latitude", "latStr", latStr)
		return nil, fmt.Errorf("failed to parse latitude: %w", err)
	}

	logger.V(2).Info("Parsing longitude", "lonStr", lonStr)
	lon, err := parseCoordinate(lonStr, logger)
	if err != nil {
		logger.Error(err, "Failed to parse longitude", "lonStr", lonStr)
		return nil, fmt.Errorf("failed to parse longitude: %w", err)
	}

	logger.V(2).Info("Checking longitude reference", "lonRef", lonRef)
	if lonRef == "West" {
		lon = -lon
	}

	gpsInfo := &GPSInfo{
		Latitude:  lat,
		Longitude: lon,
	}

	return gpsInfo, nil
}
