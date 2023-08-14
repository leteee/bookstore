package controller

import (
	"bookstore/dao"
	"bookstore/model"
	"github.com/google/uuid"
	"net/http"
)

// AddBookToCart 将图书加入到购物车
func AddBookToCart(w http.ResponseWriter, r *http.Request) {
	flag, session := dao.IsLogin(r)
	if flag {

		bookID := r.FormValue("bookId")
		book, _ := dao.GetBookByID(bookID)
		userID := session.UserID
		cart, _ := dao.GetCartByUserID(userID)
		if cart != nil {
			// 当前用户已经存在购物车
			cartItem, _ := dao.GetCartItemsByCartIDAndBookID(cart.CartID, bookID)
			if cartItem != nil {
				for _, v := range cart.CartItems {
					if v.CartItemID == cartItem.CartItemID {
						v.Count = v.Count + 1
						cartItem = v
						break
					}
				}
				dao.UpdateBookCountAndAmount(cartItem)
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
		w.Write([]byte("您刚刚将<span style='color: red'>" + book.Title + "</span>加入到了购物车中"))
	} else {
		w.Write([]byte("<span style='color: red'>请先登录！</span>"))
	}
}
