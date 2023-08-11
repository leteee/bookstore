package dao

import (
	"bookstore/model"
	"bookstore/utils"
)

// AddSession 添加session
func AddSession(s *model.Session) error {
	sqlStr := "insert into sessions values (?,?,?)"
	_, err := utils.Db.Exec(sqlStr, s.SessionID, s.Username, s.UserID)
	if err != nil {
		return err
	}
	return nil
}

// DeleteSession 添加session
func DeleteSession(sessID string) error {
	sqlStr := "delete from  sessions where session_id = ?"
	_, err := utils.Db.Exec(sqlStr, sessID)
	if err != nil {
		return err
	}
	return nil
}

// GetSessionBySessionID 查询session
func GetSessionBySessionID(sessID string) (*model.Session, error) {
	sqlStr := "select session_id, username, user_id from sessions where session_id = ?"
	row := utils.Db.QueryRow(sqlStr, sessID)
	session := &model.Session{}
	row.Scan(&session.SessionID, &session.Username, &session.UserID)
	return session, nil
}
