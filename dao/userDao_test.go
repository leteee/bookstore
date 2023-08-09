package dao

import (
	"fmt"
	"testing"
)

func Testuser(t *testing.T) {
	fmt.Println("开始测试用户")
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
