package dao

import (
	"bookstore/model"
	"bookstore/utils"
)

// AddCart 添加购物车
func AddCart(c *model.Cart) error {
	sqlStr := "insert into carts (id, total_count, total_amount, user_id) value(?,?,?,?)"
	_, err := utils.Db.Exec(sqlStr, c.CartID, c.GetTotalCount(), c.GetTotalAmount(), c.UserID)
	if err != nil {
		return err
	}
	for _, v := range c.CartItems {
		AddCartItem(v)
	}
	return nil
}

// GetCartByUserID 根据用户ID获取用户购物车
func GetCartByUserID(userID int) (*model.Cart, error) {
	sqlStr := "select id, total_count, total_amount, user_id from carts where user_id = ?"
	row := utils.Db.QueryRow(sqlStr, userID)
	cart := &model.Cart{}
	err := row.Scan(&cart.CartID, &cart.TotalCount, &cart.TotalAmount, &cart.UserID)
	if err != nil {
		return nil, err
	}
	cartItems, _ := GetCartItemsByCartID(cart.CartID)
	cart.CartItems = cartItems
	return cart, nil
}

// UpdateCart 更细购物车
func UpdateCart(cart *model.Cart) error {
	sqlStr := "update carts set total_count = ?,total_amount = ? where id = ?"
	_, err := utils.Db.Exec(sqlStr, cart.GetTotalCount(), cart.GetTotalAmount(), cart.CartID)
	if err != nil {
		return err
	}
	return nil
}
