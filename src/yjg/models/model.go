package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)


          
type Users struct {
	Id       int    `json:"id"`
	Usename  string `json:username"`
	Password  string `json:password"`	
    Sex        uint8  `json:"sex"`	
	Mobile   string `json:"mobile"`	
    Avatar string `json:"avatar"`
} 
func init() { 
	orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/test_yulongyang_?charset=utf8")
	orm.RegisterModelWithPrefix("shy_",new(Users))
	orm.RunSyncdb("default", false, true)
}
