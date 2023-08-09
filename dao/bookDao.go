package dao

import (
	"bookstore/model"
	"bookstore/utils"
)

const DefaultImgPath = "/static/img/default.jpg"

// GetBooks 查询所有图书
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

// AddBook 添加图书
func AddBook(b *model.Book) error {
	sqlStr := "insert into books (title, author, price, sales, stock, img_path) value (?, ?, ?, ?, ?, ?)"
	_, err := utils.Db.Exec(sqlStr, &b.Title, &b.Author, &b.Price, &b.Sales, &b.Stock, DefaultImgPath)
	if err != nil {
		return err
	}
	return nil
}

// DeleteBook 删除图书
func DeleteBook(bookID string) error {
	sqlStr := "delete from books where id = ?;"
	_, err := utils.Db.Exec(sqlStr, bookID)
	if err != nil {
		return err
	}
	return nil
}

// GetBookByID 根据主键获取图书
func GetBookByID(bookID string) (*model.Book, error) {
	sqlStr := "select id, title, author, price, sales, stock, img_path from books where id = ?"
	row := utils.Db.QueryRow(sqlStr, bookID)
	book := &model.Book{}
	row.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
	return book, nil
}

// UpdateBook 更新图书
func UpdateBook(b *model.Book) error {
	sqlStr := "update books set title=?, author=?, price=?, sales=?,stock=? where id=?"
	_, err := utils.Db.Exec(sqlStr, b.Title, b.Author, b.Price, b.Sales, b.Stock, b.ID)
	if err != nil {
		return err
	}
	return nil
}
