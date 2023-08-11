package dao

import (
	"bookstore/model"
	"fmt"
	"testing"
)

func TestCart(t *testing.T) {
	fmt.Println("测试购物车")
	//t.Run("测试添加购物车", testAddCart)
	t.Run("测试根据用户主键查询购物车", testGetCartByUserID)
}

func testAddCart(t *testing.T) {
	book1 := &model.Book{
		ID:    1,
		Price: 27.20,
	}
	book2 := &model.Book{
		ID:    2,
		Price: 23.00,
	}
	cartItem1 := &model.CartItem{
		Book:   book1,
		Count:  10,
		CartID: "18702761696",
	}
	cartItem2 := &model.CartItem{
		Book:   book2,
		Count:  10,
		CartID: "18702761696",
	}
	cartItems := []*model.CartItem{cartItem1, cartItem2}
	cart := &model.Cart{
		CartID:    "18702761696",
		CartItems: cartItems,
		UserID:    1,
	}
	AddCart(cart)
}

func testGetCartByUserID(t *testing.T) {
	cart, _ := GetCartByUserID(1)
	fmt.Println("查询到用户1的购物信息是：", cart)
}
