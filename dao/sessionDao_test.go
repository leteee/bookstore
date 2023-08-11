package dao

import (
	"bookstore/model"
	"fmt"
	"testing"
)

func TestSession(t *testing.T) {
	fmt.Println("开始测试用户")
	//t.Run("新增", testAddSession)
	//t.Run("查询", testGetSession)
	t.Run("删除", testDeleteSession)
}

func testAddSession(t *testing.T) {
	session := &model.Session{
		SessionID: "1870-2761-696",
		Username:  "admin",
		UserID:    1,
	}
	AddSession(session)
}

func testGetSession(t *testing.T) {
	session, _ := GetSessionBySessionID("1870-2761-696")
	fmt.Println("查询到的数据:", session)
}

func testDeleteSession(t *testing.T) {
	_ = DeleteSession("1870-2761-696")
}
