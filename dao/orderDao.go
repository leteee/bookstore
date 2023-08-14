package dao

import (
	"bookstore/model"
	"bookstore/utils"
)

func AddOrder(order *model.Order) error {
	sqlStr := "insert into orders (id, create_time, ttotal_amount, total_amount, state, user_id) values (?, ?, ?, ?, ?)"
	_, err := utils.Db.Exec(sqlStr, &order.OrderID, &order.CreateTime, &order.TotalCount, &order.TotalAmount, &order.State, &order.UserID)
	if err != nil {
		return err
	}
	return nil
}
