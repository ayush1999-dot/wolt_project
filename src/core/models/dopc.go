package models

type ResponseData struct {
	TotalPrice          float64  `json:"total_price"`
	SmallOrderSurcharge float64  `json:"small_order_surcharge"`
	CartValue           float64  `json:"cart_value"`
	Delivery            Delivery `json:"delivery"`
}

type Delivery struct {
	Fee      float64 `json:"fee"`
	Distance float64 `json:"distance"`
}
