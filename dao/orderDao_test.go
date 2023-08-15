package dao

import (
	"fmt"
	"testing"
)

func TestOrder(t *testing.T) {
	fmt.Println("开始测试订单模块")
	t.Run("测试获取所有订单", testGetOrders)
}

func testGetOrders(t *testing.T) {
	orders, _ := GetOrders()
	for _, order := range orders {
		fmt.Println("订单数据", order)
	}
}
