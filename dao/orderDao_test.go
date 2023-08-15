package dao

import (
	"fmt"
	"testing"
)

func TestOrder(t *testing.T) {
	fmt.Println("开始测试订单模块")
	//t.Run("测试获取所有订单", testGetOrders)
	t.Run("测试更细订单壮体啊", testUpdateOrderState)
}

func testGetOrders(t *testing.T) {
	orders, _ := GetOrders()
	for _, order := range orders {
		fmt.Println("订单数据", order)
	}
}

func testUpdateOrderState(t *testing.T) {
	UpdateOrderState("82157cbc-372b-4e20-a6f2-0a19b8862bf0", 2)
}
