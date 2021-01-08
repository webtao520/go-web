package  admin 


import (
	"time"
	"github.com/astaxie/beego/orm"
)




type Customer struct {
	Id        int       `orm:"column(id);auto" description:"主键"`
	Uid       string    `orm:"column(uid);size(50)" description:"用户ID"`
	Username  string    `orm:"column(username);size(255);null" description:"用户名"`
	Password  string    `orm:"column(password);size(255);null" description:"密码"`
	Nickname  string    `orm:"column(nickname);size(255);null" description:"昵称"`
	Image     string    `orm:"column(image);size(255);null" description:"头像"`
	Url       string    `orm:"column(url);size(255);null" description:"博客地址"`
	Signature string    `orm:"column(signature);size(255);null" description:"个性签名"`
	Email     string    `orm:"column(email);size(50);null" description:"邮箱"`
	Phone     string    `orm:"column(phone);size(50);null" description:"电话"`
	Wishlist  int       `orm:"column(wishlist);null" description:"收藏"`
	Review    int       `orm:"column(review);null" description:"评论"`
	Like      int       `orm:"column(like);null" description:"点赞"`
	Status    int       `orm:"column(status);null" description:"1可用，2禁用，0删除"`
	Created   time.Time `orm:"column(created);type(datetime);null" description:"创建时间"`
	Updated   time.Time `orm:"column(updated);type(datetime);null" description:"修改时间"`
}

// 自动建表
func (t *Customer) TableName()  string {
	return "customer"
}

func init (){
	orm.RegisterModel(new(Customer))
}


// 添加客户  函数
func  AddCustomer(m *Customer) (id int64,err error) {
   o:=orm.NewOrm()
   id,err=o.Insert(m)
   return	
}
 
