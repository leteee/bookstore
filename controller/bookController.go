package controller

import (
	"bookstore/dao"
	"bookstore/model"
	"net/http"
	"strconv"
	"text/template"
)

// GetBooks 获取图书
func GetBooks(w http.ResponseWriter, r *http.Request) {
	books, _ := dao.GetBooks()

	t := template.Must(template.ParseFiles("views/pages/manager/book_manager.html"))
	t.Execute(w, books)
}

// AddBook 添加图书
func AddBook(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	author := r.FormValue("title")
	price := r.FormValue("price")
	sales := r.FormValue("sales")
	stock := r.FormValue("stock")

	fPrice, _ := strconv.ParseFloat(price, 64)
	iSales, _ := strconv.ParseInt(sales, 10, 0)
	iStock, _ := strconv.ParseInt(stock, 10, 10)

	book := &model.Book{
		Title:   title,
		Author:  author,
		Price:   fPrice,
		Sales:   int(iSales),
		Stock:   int(iStock),
		ImgPath: dao.DefaultImgPath,
	}
	dao.AddBook(book)
	GetBooks(w, r)
}
