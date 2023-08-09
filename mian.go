package main

import (
	"bookstore/controller"
	"net/http"
	"text/template"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("views/index.html"))
	t.Execute(w, "")
}

func main() {
	//静态文件
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("views/static/"))))
	http.Handle("/pages/", http.StripPrefix("/pages/", http.FileServer(http.Dir("views/pages/"))))

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/index", indexHandler)

	//登录
	http.HandleFunc("/login", controller.Login)
	//注册
	http.HandleFunc("/regist", controller.Regist)
	//校验用户名
	http.HandleFunc("/checkUserName", controller.CheckUserName)
	//图书列表
	http.HandleFunc("/getBooks", controller.GetBooks)
	//添加图书
	http.HandleFunc("/addBook", controller.AddBook)
	//删除图书
	http.HandleFunc("/deleteBook", controller.DeleteBook)
	//到编辑页面
	http.HandleFunc("/toUpdateBookPage", controller.ToUpdateBookPage)
	//更新图书
	http.HandleFunc("/updateBook", controller.UpdateBook)

	http.ListenAndServe("", nil)
}

/*func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)
	http.ListenAndServe("", mux)
}*/
