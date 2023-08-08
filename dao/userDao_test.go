package dao

import (
	"fmt"
	"testing"
)

func TestUser(t *testing.T) {
	fmt.Println("测试userDao中的函数")
	t.Run("验证用户名或密码：", testSaveUser)
	t.Run("验证用户名或密码：", testRegist)
	t.Run("验证用户名或密码：", testLogin)
}

func testLogin(t *testing.T) {
	user, _ := CheckUserNameAndPassword("admin", "123456")
	fmt.Println("获取用户信息是：", user)
}

func testRegist(t *testing.T) {
	user, _ := CheckUserName("admin")
	fmt.Println("获取用户信息是：", user)
}

func testSaveUser(t *testing.T) {
	SaveUser("admin", "123456", "admin@qq.com")
}
