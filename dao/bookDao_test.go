package dao

import (
	"fmt"
	"testing"
)

func TestGetBooks(t *testing.T) {
	books, _ := GetBooks()
	for k, v := range books {
		fmt.Printf("第%v本书，%v\n", k, v)
	}
}

func TestAddBooks(t *testing.T) {
	//book := &model.Book{
	//	Title:   "三国演义",
	//	Author:  "罗贯中",
	//	Price:   88.88,
	//	Sales:   100,
	//	Stock:   100,
	//	ImgPath: DefaultImgPath,
	//}
	//AddBook(book)
}
