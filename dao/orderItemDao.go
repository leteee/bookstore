package dao

import (
	"bookstore/model"
	"bookstore/utils"
)

func AddOrderItem(o *model.OrderItem) error {
	sqlStr := "insert into order_items (order_id, count, amount, title, author, price, img_path) values (?,?,?,?,?,?,?)"
	_, err := utils.Db.Exec(sqlStr, &o.OrderID, &o.Count, &o.Amount, &o.Title, &o.Author, &o.Price, &o.ImgPath)
	if err != nil {
		return err
	}
	return nil
}
