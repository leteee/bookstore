package model

type OrderItem struct {
	OrderItemID string
	OrderID     string
	Count       int64
	Amount      float64
	Title       string
	Author      string
	Price       float64
	ImgPath     string
}
