package svc

import (
	"math"
	"testing"
)

func TestHaversine(t *testing.T) {
	tests := []struct {
		name       string
		lat1, lon1 float64
		lat2, lon2 float64
		expected   float64
	}{
		{"Same location", 52.5200, 13.4050, 52.5200, 13.4050, 0},
		{"Different location (same country)", 52.5200, 13.4050, 51.1657, 10.4515, 253.0},
		{"Large distance", 40.7128, -74.0060, 34.0522, -118.2437, 3936},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := math.Round(Haversine(tt.lat1, tt.lon1, tt.lat2, tt.lon2))
			// Compare with expected (rounded to 1 decimal point)
			if math.Round(result*10)/10 != math.Round(tt.expected*10)/10 {
				t.Errorf("expected distance %f, got %f", tt.expected, result)
			}
		})
	}

}

func TestCalculateSmallOrderSurcharge(t *testing.T) {
	tests := []struct {
		name               string
		cartValue          float64
		minimumNoSurcharge float64
		expectedSurcharge  float64
	}{
		{"No surcharge (cart >= min)", 1200, 1000, 0},
		{"With surcharge (cart < min)", 800, 1000, 200},
		{"Exact min value", 1000, 1000, 0},
		{"Large cart value", 5000, 1000, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			surcharge := CalculateSmallOrderSurcharge(tt.cartValue, tt.minimumNoSurcharge)
			if surcharge != tt.expectedSurcharge {
				t.Errorf("expected surcharge %f, got %f", tt.expectedSurcharge, surcharge)
			}
		})
	}

}
func TestCalculateDeliveryFee(t *testing.T) {
	distanceRanges := []interface{}{
		map[string]interface{}{
			"min": 0.0, "max": 500.0, "a": 0.0, "b": 0.0,
		},
		map[string]interface{}{
			"min": 500.0, "max": 1000.0, "a": 100.0, "b": 1.0,
		},
		map[string]interface{}{
			"min": 1000.0, "max": 0.0, "a": 0.0, "b": 0.0,
		},
	}

	tests := []struct {
		name          string
		distance      float64
		basePrice     float64
		expectedFee   float64
		expectedError string
	}{
		{"Delivery not possible", 1100, 199, 0, ""},
		{"Within second distance range", 750, 199, 374, ""},
		{"Within first distance range", 25.0, 150.0, 150, ""},
		{"Negative distance", -5.0, 100.0, 0, "Invalid distance"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fee := CalculateDeliveryFee(tt.distance, tt.basePrice, distanceRanges)
			if fee != tt.expectedFee {
				t.Errorf("expected fee %f, got %f", tt.expectedFee, fee)
			}
			if fee == 0 && tt.expectedError != "" {
				t.Logf("expected error: %v", tt.expectedError)
			}
		})
	}

}
