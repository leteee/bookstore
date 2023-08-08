package dao

import (
	"fmt"
	"testing"
)

func TestLogin(t *testing.T) {
	user, _ := CheckUserNameAndPassword("admin", "123456")
	fmt.Println("获取用户信息是：", user)
}

func TestRegist(t *testing.T) {
	user, _ := CheckUserName("admin")
	fmt.Println("获取用户信息是：", user)
}

func TestSaveUser(t *testing.T) {
	//SaveUser("admin", "123456", "admin@qq.com")
}
