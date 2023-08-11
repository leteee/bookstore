package model

type CartItem struct {
	CartItemID int64   //购物项主键
	Book       *Book   //购物项图书
	Count      int64   //购物项图书数量
	Amount     float64 //购物项图书的金额
	CartID     string  // 当前购物项所属购物车
}

func (c *CartItem) GetAmount() float64 {
	price := c.Book.Price
	return float64(c.Count) * price
}
