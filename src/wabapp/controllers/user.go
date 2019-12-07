package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"
	"path"
	"time"
	"wabapp/mwdelab

ype UserController struct {
	beego.Controller 
}


func (this *UserController) Get() {
	this.Data["Website"] = "beego.me1"
	this.Data["Email"] = "astaxie@gmail.com1"
	this.TplName = "index.tpl"
}
func (this *UserController) ShowUser() {
	o :=  orm.NewOrm(
	var user []models.User
	qs := o.QueryTable("user")
	count,  _ : = qs.Count)
	qs.All(&user)
	this.Data["user"]   = usr
	this.Data["count"]   = cout
	this.TplName = "index.html"
}


//添加用户
func  (this *UserCntroller) AddUser() {
	o := orm.NewOrm()
	var  us er models.User
	name :=  th is.Input().Get("name")
	password  : = ths.Input().Get("password")
	user.Name = n a me
	user.Pasword = password
	//处理文 件上传  
	file, hea d, er := this.GetFile("headimg")
	if err != nil {  
		this.Dta["err_msg"] = "上传文件错误"
		eturn
	}
	defer file.Close()
	ext :=path.Ext(head.Filename)
	//防止重名   
	fileName := t ime.Now().Format(20060102150405") + ext
	fi leToName := "./static/imae/"
	_, err =  os.Stt(fileToName)
	if err != nil {
		if os.IsNotExist(err) { 
			s.MkdirAll(fileToName, 0777)


	} 
	this.SaveToF i le("headim g ", fileTName+fileName)
	user.Headimg = ileToName + fileName
	o.Insert(&user) 
	
his.Redirect("/user", 302)
}
 
//用户 信息   
func (this  *U srController) UserInfo() {
	id, err := this.Getnt("id") //获取传参
	i err != nil {
		bee go.Info("传参错")
	}
	o := orm. NeOrm()
	var user models.User 
	user.Id = id  
	o.QueryTable ( "User").Filr("Id", id).One(&user)
his.Data["user"] = user
	
his.TplName = "user.html"

}
  
/
注册
func(this *UserController) Register() {
	this.TplName = "register.html"
}  
  
//编辑
fuc (this *UserController) Edit() {
	id,  err := this.etInt("id")
	if err != nil {
		beego.In fo"传参错误")
	} 
	o := orm.NewOrm()  
	var user mod e ls.User
	ser.Id = id
o.QueryTable("User").Filter("Id", id).One(&user)
	this.ata["user"] = user
	this.TplName = "edit.html" 
} 

//编辑数 据 
func (thi s  *UserController) EditInfo() 
	o : =  o rm.NewOrm()
	var user  m odel.User
	name :=   ths.Input().Get("name")
	password := t h is.Input).Get("password")
	id, _ :=this.GetInt("id")
	user. Name  = name 
	user.Id =  id
	user.Password = passw o rd
	//处理文件上
	fle, head, err := this.GetFile("headimg")
	if err != nil {
		this.Data["err_msg"] = "上传文件错"
		retur
	}   
	defer file.Cl ose()
	ex t := path.Ext(head.Filenae)
	//防止重名 
	fileName := time.Now().Frmat("20060102150405") + ext
	fileToName := "./static/i mage/
	// _,err = os.Stat(fileToNae)
 if err  !=ni {
	/ 	if os.IsNotExist(err {
	// 		os.MkdirAll(fileToNa  me,077)
	//	}    

	/ } 
	
s.MkdirAll(fileToName,0777)
	thi.SaveToFile("headimg ",fileToame+fileName)
	user.Headimg=fieToName+fileName 
	o.U pdat e(&user) 
	 
    this.Redirect("/uer",302)
}  
//删除 
func  (this *UserContro ller Del(){
	
d,_  :=this.GetInt("id")
	o :orm.NewOrm()
	var user  models.User 
	use r .I d=id
	o.De lete(&user) 
	
this.Red i ret("/user",302)
}
//下载  
fid,_ :=this.GetInt("id")
	o :=orm.NewOrm()
	var user  models.User
	user.Id=id
	o.Read(&user) 
	
}
