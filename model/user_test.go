package model

import (
	"fmt"
	"testing"
)

func TestMain(m *testing.M) {
	m.Run()
}

func TestUser(t *testing.T) {
	t.Run("测试添加用户：", testAddUser)
	t.Run("测试查询用户：", testGetUserById)
}

// 子测试函数
func testAddUser(t *testing.T) {
	user := &User{
		Username: "admin",
		Password: "123456",
		Email:    "admin@qq.com",
	}
	user.AddUser()
}

// 子测试函数
func testGetUserById(t *testing.T) {
	user := &User{}
	user, _ = user.GetUserById(1)
	fmt.Sprintf("查询出结果: %s %s %s", user.Username, user.Password, user.Email)
}
