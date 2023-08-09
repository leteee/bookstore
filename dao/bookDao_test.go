package dao

import (
	"bookstore/model"
	"fmt"
	"testing"
)

func TestBook(t *testing.T) {
	fmt.Println("开始测试图书模块")
	t.Run("更新图书", testUpdateBook)
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

func testDeleteBook(t *testing.T) {
	DeleteBook("32")
}

func testGetBookByID(t *testing.T) {
	book, _ := GetBookByID("12")
	fmt.Println("根据ID查询到的图书：", book)
}

func testUpdateBook(t *testing.T) {
	book := &model.Book{
		ID:     31,
		Title:  "三个女人和105个男人的故事",
		Author: "罗贯中",
		Price:  88.88,
		Sales:  10000,
		Stock:  10,
	}
	UpdateBook(book)
}
