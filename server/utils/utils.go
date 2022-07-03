package utils

import (
	"errors"
	"strconv"
	"strings"
)

func ParseLatLngQueryParam(latLng string) (lat float64, lng float64, err error) {
	if len(latLng) > 0 {
		coordinates := strings.Split(latLng, ",")
		lat, err = strconv.ParseFloat(coordinates[0], 64)
		if err != nil {
			return 0.0, 0.0, err
		}

		lng, err := strconv.ParseFloat(coordinates[1], 64)
		if err != nil {
			return 0.0, 0.0, err
		}

		return lat, lng, nil
	}

	return 0.0, 0.0, errors.New("lat lng string is empty")
}
