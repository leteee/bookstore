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

	http.ListenAndServe("", nil)
}

/*func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)
	http.ListenAndServe("", mux)
}*/
