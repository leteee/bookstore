package model

type Order struct {
	OrderID     string
	CreateTime  string
	TotalCount  int64
	TotalAmount float64
	State       int64 //订单状态 0 未发货 1 已发货 2 交易完成
	UserID      int64
}