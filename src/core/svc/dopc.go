package svc

import (
	"encoding/json"
	"fmt"
	"github.com/wolt/DOPC/src/core/models"
	"github.com/wolt/DOPC/src/utils"
	"io"
	"math"
	"net/http"
)

type DopcSvc interface {
	DopcService(venueSlug string, cartValue float64, userLat float64, userLon float64) (*models.ResponseData, error)
}

func NewDopcSvc() DopcSvc {
	return &ReceiverSvc{}
}

type ReceiverSvc struct{}

func (r *ReceiverSvc) DopcService(venueSlug string, cartValue float64, userLat float64, userLon float64) (*models.ResponseData, error) {

	responseData := &models.ResponseData{}

	staticPayload, err := http.Get(fmt.Sprintf(utils.Static, venueSlug))
	if err != nil {
		return nil, err
	}
	dynamicPayload, err := http.Get(fmt.Sprintf(utils.Dynamic, venueSlug))
	if err != nil {
		return nil, err
	}

	dynamicData, err := parseVenueJSON(dynamicPayload)
	staticData, err := parseVenueJSON(staticPayload)

	//static
	coordinates := staticData["venue_raw"].(map[string]interface{})["location"].(map[string]interface{})["coordinates"].([]interface{})
	//dynamic
	dynamicDeliverySpecs := dynamicData["venue_raw"].(map[string]interface{})["delivery_specs"].(map[string]interface{})
	deliveryPricing := dynamicDeliverySpecs["delivery_pricing"].(map[string]interface{})
	OrderMinimumNoSurcharge := dynamicDeliverySpecs["order_minimum_no_surcharge"].(float64)
	basePrice := deliveryPricing["base_price"].(float64)
	distanceRanges := deliveryPricing["distance_ranges"].([]interface{})

	//distance calculator for delivery fee
	distance := math.Round(Haversine(coordinates[1].(float64), coordinates[0].(float64), userLat, userLon))
	smallOrderSurcharg := CalculateSmallOrderSurcharge(cartValue, OrderMinimumNoSurcharge)
	totalDeliveryFee := CalculateDeliveryFee(distance, basePrice, distanceRanges)

	//totalorderfee
	totalPrice := totalDeliveryFee + smallOrderSurcharg + cartValue

	//fill response data
	responseData.TotalPrice = totalPrice
	responseData.Delivery.Fee = totalDeliveryFee
	responseData.Delivery.Distance = distance
	responseData.CartValue = cartValue
	responseData.SmallOrderSurcharge = smallOrderSurcharg
	return responseData, nil
}

func parseVenueJSON(resp *http.Response) (map[string]interface{}, error) {
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
