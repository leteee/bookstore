package controller

import (
	"bookstore/dao"
	"bookstore/model"
	"github.com/google/uuid"
	"net/http"
	"text/template"
)

// Logout 注销
func Logout(w http.ResponseWriter, r *http.Request) {
	// 获取Cookie
	cookie, _ := r.Cookie("user")
	if cookie != nil {
		cookieValue := cookie.Value
		dao.DeleteSession(cookieValue)
	}
	cookie.MaxAge = -1
	http.SetCookie(w, cookie)
	IndexHandler(w, r)
}

// Login 登录
func Login(w http.ResponseWriter, r *http.Request) {
	flag, _ := dao.IsLogin(r)
	if flag {
		IndexHandler(w, r)
	}

	username := r.FormValue("username")
	password := r.FormValue("password")
	user, _ := dao.CheckUserNameAndPassword(username, password)
	if user.ID > 0 {
		sessionID := uuid.NewString()
		session := &model.Session{
			SessionID: sessionID,
			Username:  user.Username,
			UserID:    user.ID,
		}
		dao.AddSession(session)
		cookie := &http.Cookie{
			Name:     "user",
			Value:    sessionID,
			HttpOnly: true,
		}
		http.SetCookie(w, cookie)
		t := template.Must(template.ParseFiles("views/pages/user/login_success.html"))
		t.Execute(w, user)
	} else {
		t := template.Must(template.ParseFiles("views/pages/user/login.html"))
		t.Execute(w, "用户名或密码不存在！")
	}
}

// Regist 注册
func Regist(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	email := r.FormValue("email")

	user, _ := dao.CheckUserName(username)

	if user.ID > 0 {
		//用户名被占用
		t := template.Must(template.ParseFiles("views/pages/user/regist.html"))
		t.Execute(w, "用户名已经存在！")
	} else {
		//可以使用
		dao.SaveUser(username, password, email)
		t := template.Must(template.ParseFiles("views/pages/user/regist_success.html"))
		t.Execute(w, "用户名可以使用！")
	}
}

// CheckUserName 验证用户名是否存在
func CheckUserName(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	user, _ := dao.CheckUserName(username)
	if user.ID > 0 {
		//用户名已经存在
		w.Write([]byte("用户名已经存在！"))
	} else {
		//可以使用
		w.Write([]byte("<font style='color:green'>用户名可以使用！</font>"))
	}
}
