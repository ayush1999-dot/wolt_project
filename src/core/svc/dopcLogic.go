package svc

import (
	"fmt"
	"math"
)

// CalculateSmallOrderSurcharge computes the small order surcharge.
// what does it do ? It checks if the cart value is less than the minimum value required to avoid the surcharge.
func CalculateSmallOrderSurcharge(cartValue, minimumNoSurcharge float64) float64 {
	surcharge := minimumNoSurcharge - cartValue
	if surcharge < 0 {
		return 0
	}
	return surcharge
}

// algo from internet
func Haversine(lat1, lon1, lat2, lon2 float64) float64 {
	// Convert degrees to radians
	rad := math.Pi / 180
	lat1 = lat1 * rad
	lon1 = lon1 * rad
	lat2 = lat2 * rad
	lon2 = lon2 * rad

	// Haversine formula
	dlat := lat2 - lat1
	dlon := lon2 - lon1
	a := math.Sin(dlat/2)*math.Sin(dlat/2) + math.Cos(lat1)*math.Cos(lat2)*math.Sin(dlon/2)*math.Sin(dlon/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	// Radius of the Earth in kilometers
	R := 6371.0

	// Distance in kilometers
	distance := R * c

	return distance
}

// CalculateDeliveryFee computes the delivery fee based on the distance and pre-defined pricing rules.
func CalculateDeliveryFee(distance float64, basePrice float64, distanceRanges []interface{}) float64 {

	for _, item := range distanceRanges {

		// Assert each item as a slice of maps ([][]map[string]interface{})
		mapSlice, ok := item.(map[string]interface{})
		if !ok {
			fmt.Println("Error: Item is not of type map[string]interface{}")
			continue
		}
		min, _ := mapSlice["min"].(float64)
		max, _ := mapSlice["max"].(float64)
		a, _ := mapSlice["a"].(float64)
		b, _ := mapSlice["b"].(float64)
		if distance >= min && distance < max {
			// Calculate the delivery fee
			distanceFee := math.Round(b * distance / 10)
			totalFee := a + distanceFee + basePrice
			fmt.Println(a, b)
			return totalFee
		}
	}
	return 0
}
