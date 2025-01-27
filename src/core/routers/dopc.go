package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/wolt/DOPC/src/core/handler"
	"github.com/wolt/DOPC/src/core/svc"
)

func SetupRouters(engine *gin.Engine) {

	DopcSvc := svc.NewDopcSvc()
	DopcHandler := handler.NewDopcHandler()

	engine.GET("/api/v1/delivery-order-price", DopcHandler.CalculatePrice(DopcSvc))
}
