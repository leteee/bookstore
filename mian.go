package main

import (
	"bookstore/controller"
	"net/http"
)

func main() {
	//静态文件
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("views/static/"))))
	http.Handle("/pages/", http.StripPrefix("/pages/", http.FileServer(http.Dir("views/pages/"))))

	http.HandleFunc("/", controller.IndexHandler)
	http.HandleFunc("/index", controller.IndexHandler)

	//登录
	http.HandleFunc("/login", controller.Login)
	//注销
	http.HandleFunc("/logout", controller.Logout)
	//注册
	http.HandleFunc("/regist", controller.Regist)
	//校验用户名
	http.HandleFunc("/checkUserName", controller.CheckUserName)
	//图书列表
	//http.HandleFunc("/getBooks", controller.GetBooks)
	http.HandleFunc("/getPageBooks", controller.GetPageBooks)
	//删除图书
	http.HandleFunc("/deleteBook", controller.DeleteBook)
	//到编辑页面
	http.HandleFunc("/toUpdateBookPage", controller.ToUpdateBookPage)
	//更新图书
	http.HandleFunc("/addOrUpdateBook", controller.AddOrUpdateBook)
	//添加图书到购物车
	http.HandleFunc("/addBookToCart", controller.AddBookToCart)
	//查看购物车信息
	http.HandleFunc("/getCartInfo", controller.GetCartInfo)
	//清空购物车
	http.HandleFunc("/emptyCart", controller.EmptyCart)
	//清空购物车
	http.HandleFunc("/deleteCartItem", controller.DeleteCartItem)
	//更新购物项目
	http.HandleFunc("/updateCartItem", controller.UpdateCartItem)
	//结算台
	http.HandleFunc("/checkout", controller.Checkout)
	//获取所有的订单
	http.HandleFunc("/getOrders", controller.GetOrders)
	//获取订单详情
	http.HandleFunc("/getOrderInfo", controller.GetOrderInfo)

	http.ListenAndServe("", nil)
}

/*func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)
	http.ListenAndServe("", mux)
}*/
