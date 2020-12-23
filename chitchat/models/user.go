package models 


import (
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


// create a new user,save user info the database  接受者 是指针类型
func (user *User) Create(err error){
	statement:="insert into users(uuid,name,email,password,created_at) values(?,?,?,?,?)"
	stmtin,err:=Db.Prepare(statement)
	if err !=nil {
		return 
	}
	defer stmtin.Close()
	uuid:=createUUID()
	stmtin.Exec(uuid, user.Name, user.Email, Encrypt(user.Password), time.Now())
	return
}
