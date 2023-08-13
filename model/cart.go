package model

type Cart struct {
	CartID      string
	CartItems   []*CartItem
	TotalCount  int64
	TotalAmount float64
	UserID      int
}

// GetTotalCount 计算总数量
func (c *Cart) GetTotalCount() int64 {
	var totalCount int64
	for _, v := range c.CartItems {
		totalCount = totalCount + v.Count
	}
	return totalCount
}

// GetTotalAmount 计算总金额
func (c *Cart) GetTotalAmount() float64 {
	var totalAmount float64
	for _, v := range c.CartItems {
		totalAmount = totalAmount + v.GetAmount()
	}
	return totalAmount
}
