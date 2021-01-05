package admin

import (
	"time"
	"github.com/astaxie/beego/orm"
)

const  ONLINE = 1
const  UNSALE = 2
const  DELETE = 3

var Status = map[int]string{ONLINE:"在线",UNSALE:"下架",DELETE:"删除"}

// 模型字段与数据库类型的对应 
type Article struct {
	Id        int
	Title     string
	Tag       string
	Remark    string
	Desc      string    `orm:"type(text)"`
	Html      string    `orm:"type(text)"`
	Created   time.Time `orm:"auto_now_add;type(datetime)"` // 第一次保存时才设置时间
	Updated   time.Time `orm:"auto_now;type(datetime)"` // 每次 model 保存时都会对时间自动更新
	Status    int       `orm:"default(1)"`
	Pv        int       `orm:"default(0)"`
	Review    int       `orm:"default(0)"`
	Recommend int       `orm:"default(0)"`
	Like      int       `orm:"default(0)"`
	Other     string    `orm:"type(text)"`
	Url       string
	Cover     string
	User      *User     `orm:"rel(fk)"` //  设置一对多关系
	Category  *Category `orm:"rel(one)"` // 设置一对一关系
}


// 需要在init 中注册定义的model
func init (){
	orm.RegisterModel(new(Article))
}


