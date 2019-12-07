package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"wabapp/models"
	"strconv"
	_"fmt"
	"math/rand"
	"time"
	
)
type AddressController struct{
	beego.Controller
}


func (this *AddressController) GetAddress(){
	o :=orm.NewOrm()
	var address []models.Address
	uid,_ :=this.GetInt("uid")
	o.QueryTable("Address").RelatedSel("User").Filter("User",uid).All(&address)
	
	// fmt.Println(address.User.Name)
	
	this.Data["address"]=address
	this.Data["uid"]=uid
	this.TplName="address.html"
}
//添加用户地址
func (this *AddressController) Add() {
	uid,_:=this.GetInt("uid")
	this.Data["uid"]=uid
	this.TplName="addadr.html"
}
//添加用户地址
func (this *AddressController) Addr()  {
	uid,_:=this.GetInt("uid")
	addr :=this.GetString("addr")
	o :=orm.NewOrm()
	var address models.Address
	var user models.User
	user.Id=uid
	address.Addr =addr
	address.User=&user
	o.Insert(&address)
	
	this.Redirect("/address?uid="+strconv.Itoa(uid),302)
}

func (this *AddressController) GetString(l int) string{
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	// bytes := []byte(str)
	// result := []byte{}
	// r := rand.New(rand.NewSource(time.Now().UnixNano()))
	// for i := 0; i < l; i++ {
	// 	result = append(result, bytes[r.Intn(len(bytes))])
	// }
	bytes :=[]byte(str)
	result :=[]byte{}
	r :=rand.New(rand.NewSource(time.New().UnixNano))
	for i:=0;i<l;i++{
		result =append(result,bytes[r.Intn(len(bytes))])
	}
	return string(result)

}