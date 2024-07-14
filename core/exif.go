package core

import (
	"fmt"

	"github.com/barasher/go-exiftool"
	"github.com/go-logr/logr"
)

func extractGPSInfo(imagePath string, logger logr.Logger) (string, string, string, error) {
	logger.V(1).Info("Initializing exiftool")
	et, err := exiftool.NewExiftool()
	if err != nil {
		logger.Error(err, "Failed to initialize exiftool")
		return "", "", "", fmt.Errorf("failed to initialize exiftool: %w", err)
	}
	defer et.Close()

	logger.V(1).Info("Extracting metadata", "imagePath", imagePath)
	fileInfo := et.ExtractMetadata(imagePath)

	if len(fileInfo) == 0 || fileInfo[0].Err != nil {
		logger.Error(fileInfo[0].Err, "Failed to extract metadata")
		return "", "", "", fmt.Errorf("failed to extract metadata")
	}

	latStr, ok := fileInfo[0].Fields["GPSLatitude"].(string)
	if !ok {
		logger.Error(nil, "GPSLatitude not found or not a string")
		return "", "", "", fmt.Errorf("GPSLatitude not found or not a string")
	}

	lonStr, ok := fileInfo[0].Fields["GPSLongitude"].(string)
	if !ok {
		logger.Error(nil, "GPSLongitude not found or not a string")
		return "", "", "", fmt.Errorf("GPSLongitude not found or not a string")
	}

	lonRef, ok := fileInfo[0].Fields["GPSLongitudeRef"].(string)
	if !ok {
		logger.Error(nil, "GPSLongitudeRef not found or not a string")
		return "", "", "", fmt.Errorf("GPSLongitudeRef not found or not a string")
	}

	return latStr, lonStr, lonRef, nil
}
