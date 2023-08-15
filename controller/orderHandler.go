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
					Count:   v.Count,
					Amount:  v.Amount,
					Title:   v.Book.Title,
					Author:  v.Book.Author,
					Price:   v.Book.Price,
					ImgPath: v.Book.ImgPath,
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

// GetOrders 查询所有订单
func GetOrders(w http.ResponseWriter, r *http.Request) {
	orders, _ := dao.GetOrders()
	t := template.Must(template.ParseFiles("views/pages/order/order_manager.html"))
	t.Execute(w, orders)
}

// GetOrderInfo 获取订单对应的订单项目
func GetOrderInfo(w http.ResponseWriter, r *http.Request) {
	orderID := r.FormValue("orderId")
	orderItems, _ := dao.GetOrderItemsByOrderID(orderID)
	t := template.Must(template.ParseFiles("views/pages/order/order_info.html"))
	t.Execute(w, orderItems)
}

// GetMyOrders 获取我的订单
func GetMyOrders(w http.ResponseWriter, r *http.Request) {
	flag, session := dao.IsLogin(r)
	if flag {
		orders, _ := dao.GetMyOrder(session.UserID)
		session.Orders = orders
		t := template.Must(template.ParseFiles("views/pages/order/order.html"))
		t.Execute(w, session)
	} else {
		t := template.Must(template.ParseFiles("views/pages/user/login.html"))
		t.Execute(w, "")
	}
}

// SendOrder 发货
func SendOrder(w http.ResponseWriter, r *http.Request) {
	orderId := r.FormValue("orderId")
	dao.UpdateOrderState(orderId, 1)
	GetOrders(w, r)
}

// TakeOrder 收货
func TakeOrder(w http.ResponseWriter, r *http.Request) {
	orderId := r.FormValue("orderId")
	dao.UpdateOrderState(orderId, 2)
	GetMyOrders(w, r)
}
