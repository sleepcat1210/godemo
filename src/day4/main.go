package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Class struct {
	ID   int
	Name string
	Desc string
}

//用户名:密码@tcp(主机:端口)/数据库名称?charset=utf8&parseTime=true
func main() {
	db, err := sqlx.Open(`mysql`, `root:root@tcp(127.0.0.1:3306)/news?charset=utf8&parseTime=true`)
	mod := &Class{}
	err = db.Get(mod, `select * from class limit 1`)
	fmt.Println(mod, err)
	fmt.Println("----------------------多个------------------------")
	mods := make([]Class, 0)
	err = db.Select(&mods, `select * from class`) //传入指针
	fmt.Println(mods, err)
	fmt.Println("-----------------------插入---------------------------")
	result, err := db.Exec("insert into class(`name`,`desc`) value(?,?)", `车市`, `大家好大`)
	fmt.Println(err)
	fmt.Println(result.LastInsertId())
	fmt.Println(result.RowsAffected())
	fmt.Println("-------------跟新----------------")
	result1, err1 := db.Exec("update class set `desc`=?", `你是谁1`)
	fmt.Println(err1)
	fmt.Println(result1.RowsAffected())
	fmt.Println("---------------删除--------------------------")
	result2, err2 := db.Exec("delete from class where `id`=?", `4`)
	fmt.Println(err2)
	fmt.Println(result2.RowsAffected())
	
}
