package dao

import (
	"bookstore/model"
	"bookstore/utils"
)

const DefaultImgPath = "/static/img/default.jpg"

// GetPageBooks 分页查询所有图书
func GetPageBooks(page *model.Page) (*model.Page, error) {
	sqlStr := "select count(1) from books"
	row := utils.Db.QueryRow(sqlStr)
	row.Scan(&page.TotalRecord)

	if page.TotalRecord%page.PageSize == 0 {
		page.TotalPageNo = page.TotalRecord / page.PageSize
	} else {
		page.TotalPageNo = page.TotalRecord/page.PageSize + 1
	}

	sqlStr2 := "select id, title, author, price, sales, stock, img_path from books limit ?,?"
	rows, err := utils.Db.Query(sqlStr2, (page.PageNo-1)*page.PageSize, page.PageSize)
	if err != nil {
		return nil, err
	}
	var books []*model.Book

	for rows.Next() {
		book := &model.Book{}
		rows.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
		books = append(books, book)
	}
	page.Books = books
	return page, nil
}

// GetPageBooks 分页查询所有图书
func GetPageBooksByPrice(page *model.Page, minPrice float64, macPrice float64) (*model.Page, error) {
	sqlStr := "select count(1) from books where price between ? and ?"
	row := utils.Db.QueryRow(sqlStr, minPrice, macPrice)
	row.Scan(&page.TotalRecord)

	if page.TotalRecord%page.PageSize == 0 {
		page.TotalPageNo = page.TotalRecord / page.PageSize
	} else {
		page.TotalPageNo = page.TotalRecord/page.PageSize + 1
	}

	sqlStr2 := "select id, title, author, price, sales, stock, img_path from books where price between ? and ? limit ?,?"
	rows, err := utils.Db.Query(sqlStr2, minPrice, macPrice, (page.PageNo-1)*page.PageSize, page.PageSize)
	if err != nil {
		return nil, err
	}
	var books []*model.Book

	for rows.Next() {
		book := &model.Book{}
		rows.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
		books = append(books, book)
	}
	page.Books = books
	page.MinPrice = minPrice
	page.MaxPrice = macPrice
	return page, nil
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
