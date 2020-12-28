package  models 

import "time"

type  Post struct {
	Id        int
    Uuid      string
    Body      string
    UserId    int
    ThreadId  int
    CreatedAt time.Time
}

//   Format根据layout指定的格式返回t代表的时间点的格式化文本表示。layout定义了参考时间
func (post *Post)  CreatedAtDate () string {
	 return post.CreatedAt.Format("Mon Jan 2 15:04:05 -0700 MST 2006")
}


//  获取撰写文章的用户
func (post *Post) User() (user User) {
    user = User{}
    Db.QueryRow("SELECT id, uuid, name, email, created_at FROM users WHERE id = ?", post.UserId).
        Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.CreatedAt)
    return
}

