package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/wolt/DOPC/src/core/svc"
	"net/http"
	"strconv"
)

type DopcHandler interface {
	CalculatePrice(svc svc.DopcSvc) gin.HandlerFunc
}

func NewDopcHandler() DopcHandler {
	return &Receiver{}
}

type Receiver struct{}

func (r *Receiver) CalculatePrice(svc svc.DopcSvc) gin.HandlerFunc {

	return func(c *gin.Context) {

		venueSlug := c.Query("venue_slug")
		cartValue, err := strconv.ParseFloat(c.Query("cart_value"), 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid cart value"})
			return
		}
		userLatStr := c.Query("user_lat")
		userLonStr := c.Query("user_lon")

		userLat, err := strconv.ParseFloat(userLatStr, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user_lat parameter"})
			return
		}

		userLon, err := strconv.ParseFloat(userLonStr, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user_lon parameter"})
			return
		}

		resp, err := svc.DopcService(venueSlug, cartValue, userLat, userLon)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, resp)
	}

}
