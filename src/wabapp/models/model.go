package models

import (	
	_"github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	
)

type User struct {
	Id       int
	Name     string
	Password string
	Headimg	string
	Address   []*Address   `orm:"reverse(many)"`
}

type Address struct {
	Id	int	`orm:"pk;auto"`
	Addr	string
	User		*User `orm:"rel(fk)"` 
}
func init() {
	orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/news?charset=utf8")
	orm.RegisterModel(new(User),new(Address))
	orm.RunSyncdb("default", false, true)
}
