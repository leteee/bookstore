package model

import (
	"bookstore/utils"
	"fmt"
)

type User struct {
	ID       int
	Username string
	Password string
	Email    string
}

func (user *User) AddUser() error {
	// 写sql语句
	sqlStr := "insert into users(username , password , email) values(?,?,?)"

	stmt, err := utils.Db.Prepare(sqlStr)

	if err != nil {
		fmt.Println("预编译出现异常：", err)
		return err
	}

	_, erro := stmt.Exec(user.Username, user.Password, user.Email)
	if erro != nil {
		fmt.Println("执行出现异常：", erro)
		return erro
	}
	return nil
}

func (user *User) GetUserById(userId int) (*User, error) {
	// 写sql语句
	sqlStr := "select id , username , password , email from users where id = ?"

	stmt, err := utils.Db.Prepare(sqlStr)

	if err != nil {
		fmt.Println("预编译出现异常：", err)
		return nil, err
	}

	var username, password, email string

	// 执行sql
	row := stmt.QueryRow(userId)

	err = row.Scan(&userId, &username, &password, &email)
	if err != nil {
		fmt.Println("执行出现异常：", err)
		return nil, err
	}

	u := &User{
		ID:       userId,
		Username: username,
		Password: password,
		Email:    email,
	}
	return u, nil
}
