package dao

import (
	"bookstore/model"
	"fmt"
	"testing"
)

func TestBook(t *testing.T) {
	fmt.Println("开始测试图书模块")
	t.Run("测试删除", deleteBook)
	t.Run("查询一个", getBookByID)
}

func testGetBooks(t *testing.T) {
	books, _ := GetBooks()
	for k, v := range books {
		fmt.Printf("第%v本书，%v\n", k, v)
	}
}

func testAddBooks(t *testing.T) {
	book := &model.Book{
		Title:   "三国演义",
		Author:  "罗贯中",
		Price:   88.88,
		Sales:   100,
		Stock:   100,
		ImgPath: DefaultImgPath,
	}
	AddBook(book)
}

func deleteBook(t *testing.T) {
	DeleteBook("32")
}

func getBookByID(t *testing.T) {
	book, _ := GetBookByID("12")
	fmt.Println("根据ID查询到的图书：", book)
}
