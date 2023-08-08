package dao

import (
	"bookstore/model"
	"bookstore/utils"
)

const DefaultImgPath = "/static/img/default.jpg"

func GetBooks() ([]*model.Book, error) {
	sqlStr := "select id, title, author, price, sales, stock, img_path from books"
	rows, err := utils.Db.Query(sqlStr)
	if err != nil {
		return nil, err
	}
	var books []*model.Book

	for rows.Next() {
		book := &model.Book{}
		rows.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
		books = append(books, book)
	}
	return books, nil
}

func AddBook(b *model.Book) error {
	sqlStr := "insert into books (title, author, price, sales, stock, img_path) value (?, ?, ?, ?, ?, ?)"
	_, err := utils.Db.Exec(sqlStr, &b.Title, &b.Author, &b.Price, &b.Sales, &b.Stock, DefaultImgPath)
	if err != nil {
		return err
	}
	return nil
}
