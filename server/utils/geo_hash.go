package utils

import (
	"math"

	"github.com/mmcloughlin/geohash"
)

// Reference: www.elastic.co/guide/en/elasticsearch/reference/current/search-aggregations-bucket-geohashgrid-aggregation.html

// unit in meter
var cellWidthPerPrecision = []float64{5009400.0, 1252300.0, 156500.0, 39100.0, 4900.0, 1200.0, 152.9, 38.2, 4.8, 1.2, 0.149, 0.037}

// unit in meter
var cellHeightPerPrecision = []float64{4992600.0, 624100.0, 156000.0, 19500.0, 4900.0, 609.4, 152.4, 19.0, 4.8, 0.595, 0.149, 0.019}

func summarizeGeoHash(geoHashes []string, minPrecision int, maxPrecision int) {

}

func degreeToRadian(degree float64) float64 {
	return degree * (math.Pi / 180)
}

func radianToDegree(radian float64) float64 {
	return radian * (180 / math.Pi)
}

func latLngDistance(lat1 float64, lng1 float64, lat2 float64, lng2 float64) float64 {
	var earthRadius = 6371000 // radius of the earth in m
	var radLat = degreeToRadian(lat2 - lat1)
	var radLng = degreeToRadian(lng2 - lng1)
	var a = math.Sin(radLat/2)*math.Sin(radLat/2) +
		math.Cos(degreeToRadian(lat1))*math.Cos(degreeToRadian(lat2))*
			math.Sin(radLng/2)*math.Sin(radLng/2)

	return float64(earthRadius) * 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
}

// Reference: https://stackoverflow.com/questions/7477003/calculating-new-longitude-latitude-from-old-n-meters
func addXYToLatLng(height float64, width float64, latitude float64, longitude float64) (lat, lng float64) {
	var earthRadius = 6371000.0 // radius of the earth in m

	latDiff := radianToDegree(float64(height) / earthRadius)
	lngDiff := radianToDegree(float64(width)/earthRadius) / math.Cos(latitude*math.Pi/180)

	return latitude + latDiff, longitude + lngDiff
}

func findGeoHash(latLngs [][]float64, precision int) []string {
	// find a rectangle enclosing this shape
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

	// generate geohash with precision which lie inside this box
	latSegments := int(math.Ceil(latLngDistance(boundingBox.MinLat, boundingBox.MinLng, boundingBox.MaxLat, boundingBox.MinLng) / cellHeightPerPrecision[precision-1]))
	lngSegments := int(math.Ceil(latLngDistance(boundingBox.MinLat, boundingBox.MinLng, boundingBox.MinLat, boundingBox.MaxLng) / cellWidthPerPrecision[precision-1]))

	curHeight := 0.0
	curWidth := 0.0
	geoHashes := make(map[string]struct{})
	heightFactor := cellHeightPerPrecision[precision-1] / 2
	widthFactor := cellWidthPerPrecision[precision-1] / 2
	for h := 0; h <= latSegments; h++ {
		curHeight = heightFactor * float64(h)
		for w := 0; w <= lngSegments; w++ {
			curWidth = widthFactor * float64(w)
			newLat, newLng := addXYToLatLng(-1*curHeight, curWidth, boundingBox.MaxLat, boundingBox.MinLng)
			geoHashes[geohash.EncodeWithPrecision(newLat, newLng, uint(precision))] = struct{}{}
		}
	}

	// TODO:filter out all those geohash which doesnt lie inside this shape
	index := 0
	hashes := make([]string, len(geoHashes))
	for geoHash := range geoHashes {
		hashes[index] = geoHash
		index++
	}

	return hashes
}
