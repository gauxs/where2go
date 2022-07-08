package db

import (
	"math"
	"math/rand"
	"time"

	"github.com/mmcloughlin/geohash"
)

var geoHash map[string]string

func init() {
	geoHash = make(map[string]string)
}

/*
	Input:
	2 pair of (lat, lng)
	from and to date
	count - number of locations to generate
*/
func generateData(latLngs [][]float64, count int) {
	minLat := math.MaxFloat64
	maxLat := math.SmallestNonzeroFloat64
	minLng := math.MaxFloat64
	maxLng := math.SmallestNonzeroFloat64
	for _, latLng := range latLngs {
		minLat = math.Min(minLat, latLng[0])
		maxLat = math.Max(maxLat, latLng[0])

		minLng = math.Min(minLng, latLng[1])
		maxLng = math.Max(maxLng, latLng[1])
	}

	boundingBox := geohash.Box{
		MinLat: minLat,
		MaxLat: maxLat,
		MinLng: minLng,
		MaxLng: maxLng,
	}

	precision := 8
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < count; i++ {
		randLat := boundingBox.MinLat + rand.Float64()*(boundingBox.MaxLat-boundingBox.MinLat)
		randLng := boundingBox.MinLng + rand.Float64()*(boundingBox.MaxLng-boundingBox.MinLng)

		geohash.EncodeWithPrecision(randLat, randLng, uint(precision))

		// static data

		// non-static data
		// range over [from, to] date + random event

	}
}
