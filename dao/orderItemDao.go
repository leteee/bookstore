package dao

import (
	"bookstore/model"
	"bookstore/utils"
	"fmt"
)

func AddOrderItem(o *model.OrderItem) error {
	sqlStr := "insert into order_items (order_id, count, amount, title, author, price, img_path) values (?,?,?,?,?,?,?)"
	_, err := utils.Db.Exec(sqlStr, &o.OrderID, &o.Count, &o.Amount, &o.Title, &o.Author, &o.Price, &o.ImgPath)
	if err != nil {
		return err
	}
	return nil
}

func GetOrderItemsByOrderID(orderID string) ([]*model.OrderItem, error) {
	sqlStr := "select id, order_id, count, amount, title, author, price, img_path from order_items where order_id = ?"
	rows, err := utils.Db.Query(sqlStr, orderID)
	if err != nil {
		return nil, err
	}
	var orderItems []*model.OrderItem
	for rows.Next() {
		orderItem := &model.OrderItem{}
		err := rows.Scan(&orderItem.OrderItemID, &orderItem.OrderID, &orderItem.Count, &orderItem.Amount, &orderItem.Title, &orderItem.Author, &orderItem.Price, &orderItem.ImgPath)
		if err != nil {
			fmt.Println(err)
		}
		orderItems = append(orderItems, orderItem)
	}
	return orderItems, nil
}
