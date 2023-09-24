package helper

import (
	"math"
	"strconv"
)

type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longtitude"`
}

const (
	earthRadiusKm = 6371 // Earth's radius in kilometers
)

func degreeToRadian(degree float64) float64 {
	return degree * math.Pi / 180
}

func HaversineDistance(loc1, loc2 Location) float64 {
	lat1 := degreeToRadian(loc1.Latitude)
	lon1 := degreeToRadian(loc1.Longitude)
	lat2 := degreeToRadian(loc2.Latitude)
	lon2 := degreeToRadian(loc2.Longitude)

	dLat := lat2 - lat1
	dLon := lon2 - lon1

	a := math.Pow(math.Sin(dLat/2), 2) + math.Cos(lat1)*math.Cos(lat2)*math.Pow(math.Sin(dLon/2), 2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	distance := earthRadiusKm * c
	//log.Println(distance)
	return distance
}

func Convert(l string) float64 {
	// Convert l1 to float64
	value64_1, _ := strconv.ParseFloat(l, 64)
	return value64_1
}

func CheckDistance(loc Location, l1, l2 float64) bool {
	var loc2 Location
	loc2.Latitude = l1
	loc2.Longitude = l2
	if HaversineDistance(loc, loc2) <= 25.0 {
		return true
	} else {
		return false
	}
}
