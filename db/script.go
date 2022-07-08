package db

import (
	"math"
	"math/rand"
	"time"

	"github.com/mmcloughlin/geohash"
)

var actorNames = []string{
	"Sharukh Khan", "Akshay Kumar", "Salman Khan", "Ajay Devgn",
	"Aamir Khan", "Hrithik Roshan", "Shahid Kapoor", "Tiger Shroff",
	"Ranveer Singh", "Varun Dhawan", "Saif Ali Khan"}

type PlaceAttribute int

const (
	ActorName PlaceAttribute = iota
	HotelBooking
	NumberOfVisitors
)

type StaticGeo struct {
	GeoHash        string
	Attribute      PlaceAttribute
	AttributeValue interface{}
}

type TimeGeo struct {
	GeoHash        string
	Date           time.Time
	Attribute      PlaceAttribute // will be same for each row
	AttributeValue interface{}
}

type Geo struct {
	GeoHash   string
	Latitude  float64
	Longitude float64
	Type      string // Hotel | Amusement | Religious | Nature
	Name      string // pool of pre-defined names
	Static    []StaticGeo
	Time      []TimeGeo
}

var geoHash map[string]Geo

func init() {
	geoHash = make(map[string]Geo)
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

		hash := geohash.EncodeWithPrecision(randLat, randLng, uint(precision))
		if _, ok := geoHash[hash]; ok {
			// static data
			// either a hotel or a site
			// actor visited with less probability
			// static := StaticGeo{
			// 	GeoHash:        hash,
			// 	Attribute:      ActorName, // HotelName | ActorName
			// 	AttributeValue: "Sharukh Khan",
			// }

			// add actor
			// non-static data
			// range over [from, to] date + random event

			// geo := Geo{
			// 	GeoHash:   hash,
			// 	Latitude:  randLat,
			// 	Longitude: randLng,
			// }
		} else {
			i--
			continue
		}
	}
}
