package core

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/go-logr/logr"
)

func parseCoordinate(coord string, logger logr.Logger) (float64, error) {
	logger.V(3).Info("Parsing coordinate", "coord", coord)
	parts := strings.Fields(coord)
	logger.V(3).Info("Split coordinate into parts", "parts", parts, "count", len(parts))

	if len(parts) != 4 && len(parts) != 5 {
		return 0, fmt.Errorf("invalid coordinate format: expected 4 or 5 parts, got %d", len(parts))
	}

	deg, err := strconv.ParseFloat(strings.Trim(parts[0], "deg"), 64)
	if err != nil {
		logger.Error(err, "Failed to parse degrees", "part", parts[0])
		return 0, fmt.Errorf("failed to parse degrees: %w", err)
	}
	logger.V(3).Info("Parsed degrees", "deg", deg)

	min, err := strconv.ParseFloat(strings.Trim(parts[2], "'"), 64)
	if err != nil {
		logger.Error(err, "Failed to parse minutes", "part", parts[2])
		return 0, fmt.Errorf("failed to parse minutes: %w", err)
	}
	logger.V(3).Info("Parsed minutes", "min", min)

	sec, err := strconv.ParseFloat(strings.Trim(parts[3], "\""), 64)
	if err != nil {
		logger.Error(err, "Failed to parse seconds", "part", parts[3])
		return 0, fmt.Errorf("failed to parse seconds: %w", err)
	}
	logger.V(3).Info("Parsed seconds", "sec", sec)

	result := deg + min/60 + sec/3600
	logger.V(3).Info("Calculated final coordinate", "result", result)
	return result, nil
}
