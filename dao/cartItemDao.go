package dao

import (
	"bookstore/model"
	"bookstore/utils"
)

// AddCartItem Tina集购物车项目
func AddCartItem(c *model.CartItem) error {
	sqlStr := "insert into cart_items (count, amount, book_id, cart_id) value(?,?,?,?)"
	_, err := utils.Db.Exec(sqlStr, c.Count, c.GetAmount(), c.Book.ID, c.CartID)
	if err != nil {
		return err
	}
	return nil
}

// GetCartItemByBookID 根据图书的id获取对应的购物项目
func GetCartItemByBookID(bookId int) (*model.CartItem, error) {
	sqlStr := "select id, count, amount, cart_id from cart_items where book_id = ?"
	row := utils.Db.QueryRow(sqlStr, bookId)
	cartItem := &model.CartItem{}
	err := row.Scan(&cartItem.CartItemID, &cartItem.Count, &cartItem.Amount, &cartItem.CartID)
	if err != nil {
		return nil, err
	}
	return cartItem, nil
}

// GetCartItemsByCartID 根据购物车的id获取购物车中所有的购物项
func GetCartItemsByCartID(cartId string) ([]*model.CartItem, error) {
	sqlStr := "select id, count, amount, cart_id from cart_items where cart_id = ?"
	rows, err := utils.Db.Query(sqlStr, cartId)
	if err != nil {
		return nil, err
	}

	var cartItems []*model.CartItem
	for rows.Next() {
		cartItem := &model.CartItem{}
		err := rows.Scan(&cartItem.CartItemID, &cartItem.Count, &cartItem.Amount, &cartItem.CartID)
		if err != nil {
			return nil, err
		}
		cartItems = append(cartItems, cartItem)
	}
	return cartItems, nil
}
