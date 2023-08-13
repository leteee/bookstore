package controller

import (
	"bookstore/dao"
	"bookstore/model"
	"github.com/google/uuid"
	"net/http"
)

// Logout 注销
func AddBookToCart(w http.ResponseWriter, r *http.Request) {
	bookID := r.FormValue("bookId")
	book, _ := dao.GetBookByID(bookID)
	_, session := dao.IsLogin(r)
	userID := session.UserID
	cart, _ := dao.GetCartByUserID(userID)
	if cart != nil {
		// 当前用户已经存在购物车
		cartItem, _ := dao.GetCartItemsByCartIDAndBookID(cart.CartID, bookID)
		if cartItem != nil {
			cartItem.Count = cartItem.Count + 1
		} else {
			// 还没有该图书
			cartItem = &model.CartItem{
				Book:   book,
				Count:  1,
				Amount: book.Price,
				CartID: cart.CartID,
			}
			// 将购物车项目添加到当前cart中
			cart.CartItems = append(cart.CartItems, cartItem)
			dao.AddCartItem(cartItem)
		}
		dao.UpdateCart(cart)
	} else {
		//创建购物车
		cartID := uuid.NewString()
		cart = &model.Cart{
			CartID: cartID,
			UserID: userID,
		}
		cart.CartItems = []*model.CartItem{
			{
				Book:   book,
				Count:  1,
				Amount: book.Price,
				CartID: cartID,
			},
		}
		dao.AddCart(cart)
	}
}
