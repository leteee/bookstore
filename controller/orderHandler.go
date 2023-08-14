package controller

import (
	"bookstore/dao"
	"bookstore/model"
	"github.com/google/uuid"
	"html/template"
	"net/http"
	"time"
)

// Checkout 结算
func Checkout(w http.ResponseWriter, r *http.Request) {
	flag, session := dao.IsLogin(r)
	if flag {
		cart, err := dao.GetCartByUserID(session.UserID)
		if err == nil && len(cart.CartItems) > 0 {
			userID := session.UserID
			orderID := uuid.NewString()
			order := &model.Order{
				OrderID:     orderID,
				CreateTime:  time.Now().Format("2006-01-02 15:04:05"),
				TotalCount:  cart.TotalCount,
				TotalAmount: cart.TotalAmount,
				State:       0,
				UserID:      int64(int(userID)),
			}
			dao.AddOrder(order)
			cartItems := cart.CartItems
			for _, v := range cartItems {
				orderItem := &model.OrderItem{
					OrderID: orderID,
					Book:    v.Book,
					Count:   v.Count,
					Amount:  v.Amount,
				}
				dao.AddOrderItem(orderItem)
				book := v.Book
				book.Stock = book.Stock - int(v.Count)
				book.Sales = book.Sales + int(v.Count)
				dao.UpdateBook(book)
			}
			dao.DeleteCartByID(cart.CartID)
			session.Order = order
			t := template.Must(template.ParseFiles("views/pages/cart/checkout.html"))
			t.Execute(w, session)
		} else {
			t := template.Must(template.ParseFiles("views/pages/cart/cart.html"))
			t.Execute(w, session)
		}
	} else {
		t := template.Must(template.ParseFiles("views/pages/user/login.html"))
		t.Execute(w, "")
	}
}
