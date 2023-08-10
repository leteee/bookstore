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

// GetPageBooks 分页获取图书
func GetPageBooks(w http.ResponseWriter, r *http.Request) {
	pageNo := r.FormValue("pageNo")
	if pageNo == "" {
		pageNo = "1"
	}
	iPageNo, _ := strconv.ParseInt(pageNo, 10, 64)
	page := &model.Page{
		PageNo:   iPageNo,
		PageSize: 4,
	}
	page, _ = dao.GetPageBooks(page)
	t := template.Must(template.ParseFiles("views/pages/manager/book_manager.html"))
	t.Execute(w, page)
}

// DeleteBook 删除图书
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	bookID := r.FormValue("bookId")
	dao.DeleteBook(bookID)
	GetBooks(w, r)
}

// ToUpdateBookPage 修改图书
func ToUpdateBookPage(w http.ResponseWriter, r *http.Request) {
	bookID := r.FormValue("bookId")
	book, _ := dao.GetBookByID(bookID)

	if book.ID > 0 {
		t := template.Must(template.ParseFiles("views/pages/manager/book_edit.html"))
		t.Execute(w, book)
	} else {
		t := template.Must(template.ParseFiles("views/pages/manager/book_edit.html"))
		t.Execute(w, "")
	}
}

// AddOrUpdateBook 新增或更新图书
func AddOrUpdateBook(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("bookId")
	title := r.FormValue("title")
	author := r.FormValue("author")
	price := r.FormValue("price")
	sales := r.FormValue("sales")
	stock := r.FormValue("stock")

	iId, _ := strconv.ParseInt(id, 10, 0)
	fPrice, _ := strconv.ParseFloat(price, 64)
	iSales, _ := strconv.ParseInt(sales, 10, 0)
	iStock, _ := strconv.ParseInt(stock, 10, 0)

	book := &model.Book{
		ID:     int(iId),
		Title:  title,
		Author: author,
		Price:  fPrice,
		Sales:  int(iSales),
		Stock:  int(iStock),
	}
	if book.ID > 0 {
		dao.UpdateBook(book)
	} else {
		dao.AddBook(book)
	}
	GetBooks(w, r)
}
