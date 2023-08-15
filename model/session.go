package model

type Session struct {
	SessionID string
	Username  string
	UserID    int
	Cart      *Cart
	Order     *Order
	Orders    []*Order
}
