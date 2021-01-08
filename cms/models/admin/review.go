package admin

import (
	_"fmt"
	"github.com/astaxie/beego/orm"
	"time"
)

// 评论 
type Review struct {
	Id        int
	Name      string
	Review    string    `orm:"size(500)"`
	Reply     string    `orm:"size(500)"`
	Site      string    `orm:"size(500)"`
	Created   time.Time `orm:"auto_now_add;type(datetime)"`
	Updated   time.Time `orm:"auto_now;type(datetime)"`
	Status    int       `orm:"default(1)"`
	ArticleId int
	Customer  *Customer `orm:"rel(fk)"`
	Like      int
	Star      int
}

// 需要在init中注册定义model
func init (){
	orm.RegisterModel(new(Review))
}