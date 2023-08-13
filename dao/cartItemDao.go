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
func GetCartItemByBookID(bookID string) (*model.CartItem, error) {
	sqlStr := "select id, count, amount, book_id, cart_id from cart_items where book_id = ?"
	row := utils.Db.QueryRow(sqlStr, bookID)
	cartItem := &model.CartItem{}
	var bookId string
	err := row.Scan(&cartItem.CartItemID, &cartItem.Count, &cartItem.Amount, &bookId, &cartItem.CartID)
	if err != nil {
		return nil, err
	}
	cartItem.Book, _ = GetBookByID(bookId)
	return cartItem, nil
}

// GetCartItemsByCartID 根据购物车的id获取购物车中所有的购物项
func GetCartItemsByCartID(cartID string) ([]*model.CartItem, error) {
	sqlStr := "select id, count, amount, book_id, cart_id from cart_items where cart_id = ?"
	rows, err := utils.Db.Query(sqlStr, cartID)
	if err != nil {
		return nil, err
	}
	var cartItems []*model.CartItem
	for rows.Next() {
		var bookId string
		cartItem := &model.CartItem{}
		err := rows.Scan(&cartItem.CartItemID, &cartItem.Count, &cartItem.Amount, &bookId, &cartItem.CartID)
		if err != nil {
			return nil, err
		}
		cartItem.Book, _ = GetBookByID(bookId)
		cartItems = append(cartItems, cartItem)
	}
	return cartItems, nil
}

// GetCartItemsByCartIDAndBookID 根据购物车主键和书籍主键获取购物车中所有的购物项
func GetCartItemsByCartIDAndBookID(cartID, bookID string) (*model.CartItem, error) {
	sqlStr := "select id, count, amount, book_id, cart_id from cart_items where cart_id = ? and book_id = ?"
	rows := utils.Db.QueryRow(sqlStr, cartID, bookID)
	cartItem := &model.CartItem{}
	var bookId string
	err := rows.Scan(&cartItem.CartItemID, &cartItem.Count, &cartItem.Amount, &bookId, &cartItem.CartID)
	if err != nil {
		return nil, err
	}
	cartItem.Book, _ = GetBookByID(bookId)
	return cartItem, nil
}
