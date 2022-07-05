package utils

import (
	"testing"
)

func TestGeoHash(t *testing.T) {
	var latLngs = make([][]float64, 4)
	latLngs[0] = []float64{36.21, 65.45}
	latLngs[1] = []float64{36.21, 100.0}
	latLngs[2] = []float64{6.58, 100.0}
	latLngs[3] = []float64{6.58, 67.67}

	geoHashesOfIndia := findGeoHash(latLngs, 2)
	for _, hash := range geoHashesOfIndia {
		t.Log(hash)
	}
}
