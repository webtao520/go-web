package models

import (
	_"fmt"
	"time"
)

type User struct {
	Id        int
    Uuid      string
    Name      string
    Email     string
    Password  string
    CreatedAt time.Time // time包提供了时间的显示和测量用的函数。日历的计算采用的是公历。
}


// Create a new user, save user info into the database
func (user *User) Create() (err error) {
    // Postgres does not automatically return the last insert id, because it would be wrong to assume
    // you're always using a sequence.You need to use the RETURNING keyword in your insert to get this
    // information from postgres.
    statement := "insert into users (uuid, name, email, password, created_at) values (?, ?, ?, ?, ?)"
    stmtin, err := Db.Prepare(statement)
    if err != nil {
        return
    }
    defer stmtin.Close()

    uuid := createUUID()
    stmtin.Exec(uuid, user.Name, user.Email, Encrypt(user.Password), time.Now())

	/*
    stmtout, err := Db.Prepare("select id, uuid, created_at from users where uuid = ?")
    if err != nil {
        return
    }
    defer stmtout.Close()
    // use QueryRow to return a row and scan the returned id into the User struct
	err = stmtout.QueryRow(uuid).Scan(&user.Id, &user.Uuid, &user.CreatedAt)
	*/
    return
}


// get a single user given the email 
func  UserByEmail(email string)(user User, err error){
	//  初始化分配内存
	// func（rs * Rows）Scan（dest ... interface {}） 错误  扫描将当前行中的列复制到dest指向的值中。dest中的值数必须与“行”中的列数相同。
    user = User{}
    err = Db.QueryRow("SELECT id, uuid, name, email, password, created_at FROM users WHERE email = ?", email).
        Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
    return
}

// create a  new session for an existing user
func (user *User) CreateSession() (session Session,err error){
	statement:="insert into sessions (uuid,email,user_id,created_at)  values (?, ?, ?, ?)"
	stmtin,err:=Db.Prepare(statement)
	if err != nil {
		return
	}
	defer  stmtin.Close()
	uuid:=createUUID()
	// func (db *DB) Exec(query string, args ...interface{}) (Result, error)
	// Exec执行查询而不返回任何行。args用于查询中的任何占位符参数。
	stmtin.Exec(uuid, user.Email, user.Id, time.Now())

	stmtout,err:=Db.Prepare("select id, uuid, email, user_id, created_at from sessions where uuid = ?")
	if err!=nil {
		return
	}
	defer stmtout.Close()
	err=stmtout.QueryRow(uuid).Scan(&session.Id, &session.Uuid, &session.Email, &session.UserId, &session.CreatedAt)
	return

}

 // 创建一个新的主题
 func (user *User) CreateThread(topic string) (conv Thread, err error) {
    statement := "insert into threads (uuid, topic, user_id, created_at) values (?, ?, ?, ?)"
    stmtin, err := Db.Prepare(statement)
    if err != nil {
        return
    }
    defer stmtin.Close()

    uuid := createUUID()
    stmtin.Exec(uuid, topic, user.Id, time.Now())

    stmtout, err := Db.Prepare("select id, uuid, topic, user_id, created_at from threads where uuid = ?")
    if err != nil {
        return
    }
    defer stmtout.Close()

    // use QueryRow to return a row and scan the returned id into the Session struct
    err = stmtout.QueryRow(uuid).Scan(&conv.Id, &conv.Uuid, &conv.Topic, &conv.UserId, &conv.CreatedAt)
    return
}





