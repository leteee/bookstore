package dao

import (
	"fmt"
	"testing"
)

func TestOrderItem(t *testing.T) {
	fmt.Println("开始测试订单模块")
	t.Run("测试获取订单项目", testGetOrderItemsByOrderID)
}

func testGetOrderItemsByOrderID(t *testing.T) {
	orderItems, _ := GetOrderItemsByOrderID("82157cbc-372b-4e20-a6f2-0a19b8862bf0")
	for _, orderItem := range orderItems {
		fmt.Println("订单详情数据", orderItem)
	}
}
