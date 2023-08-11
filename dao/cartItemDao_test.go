package dao

import (
	"fmt"
	"testing"
)

func TestCartItem(t *testing.T) {
	fmt.Println("测试购物项目")
	//t.Run("测试根据图书的id获取对应的购物项目", testGetCartItemByBookID)
	//t.Run("测试根据购物车的id获取购物车中所有的购物项", testGetCartItemsByCartID)
}

func testGetCartItemByBookID(t *testing.T) {
	cartItem, _ := GetCartItemByBookID(1)
	fmt.Println("图书id=1的购物项信息是：", cartItem)
}

func testGetCartItemsByCartID(t *testing.T) {
	cartItems, _ := GetCartItemsByCartID("18702761696")
	for _, v := range cartItems {

		fmt.Println("购物车18702761696的购物项信息是：", v)
	}
}
