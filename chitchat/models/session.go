package models

import "time"

// session 结构体

type Session struct {
	Id        int
	Uuid      string
	Email     string
	UserId    int
	CreatedAt time.Time
}

// 从数据库中删除会话
func (session *Session) DeleteByUUID() (err error) {
	statement := "delete from sessions where uuid = ?"
	// Prepare为以后的查询或执行创建一个准备好的语句。可以从返回的语句中并发运行多个查询或执行。当不再需要该语句时，调用方必须调用该语句的Close方法。
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	// Exec执行查询而不返回任何行。args用于查询中的任何占位符参数。
	_, err = stmt.Exec(session.Uuid)
	return
}

// 检查数据库中的session是否有效
func (session *Session) Check() (valid bool, err error) {
	err = Db.QueryRow("select id, uuid, email, user_id, created_at from sessions where uuid = ?", session.Uuid).Scan(
		&session.Id, &session.Uuid, &session.Email, &session.UserId, &session.CreatedAt)
	if err != nil {
		valid = false
		return
	}
	if session.Id != 0 {
		valid = true
	}
	return
}

// 从回话中 获取session信息，查询用户信息
func (session *Session) User() (user User, err error) {
    user = User{}
    err = Db.QueryRow("SELECT id, uuid, name, email, created_at FROM users WHERE id = ?", session.UserId).
        Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.CreatedAt)
    return
}