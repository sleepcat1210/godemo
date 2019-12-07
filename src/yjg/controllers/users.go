package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"yjg/models"
	"crypto/md5"
	_"encoding/hex"
	"encoding/base64"
	"time"
	"strconv"
	"fmt"
	"path"
	"os"
	
	
)

type UsersController struct {
	beego.Controller
}

func (this *UsersController) Get() {
	o :=orm.NewOrm()
	var user []models.Users
	var result =make(map[string]interface{})
	o.QueryTable("shy_users").All(&user)
	result["code"]=200
	result["msg"]="获取数据成功！"
	result["data"]=user
	this.Data["json"]=&result
	this.ServeJSON()
}
func (this *UsersController) Login(){
	o :=orm.NewOrm()
	var user models.Users
	mobile :=this.GetString("mobile")
	password :=this.GetString("password")
	var result =make(map[string]interface{})
	if mobile =="" || password==""{
		result["code"]=201
		result["msg"]="用户和密码不能为空！"
		this.Data["json"]=&result
	    this.ServeJSON()
	}
	user.Mobile =mobile
	err :=o.Read(&user,"Mobile")
	if err != nil{
		result["code"]=201
		result["msg"]="账户不存在！"
		this.Data["json"]=&result
	    this.ServeJSON()
	}
	o.QueryTable("shy_users").Filter("mobile",mobile).One(&user)
	if user.Password == password{
		result["code"]=200
		result["msg"]="登录成功！"
		result["data"]=user
		this.Data["json"]=&result
	    this.ServeJSON()
	}else{
		result["code"]=201
		result["msg"]="账户或密码不正确！"
		result["data"]=user		
	
		this.Data["json"]=&result
	    this.ServeJSON()
	}	
	
}

//上传头像
func (this *UsersController) UploadImg(){
	uid,_:=this.GetInt("uid")
	file, head, err :=this.GetFile("headimg")
	var result =make(map[string]interface{})
	if err !=nil{
		result["code"]=100
		result["msg"]="上传文件错误！"
		this.Data["json"]=&result
		this.ServeJSON()
	}
	defer file.Close()
	ext :=path.Ext(head.Filename)
	fileName :=time.Now().Format("20060102150405") + ext
	fileToName :="./static/img/"
	os.MkdirAll(fileToName,0777)
	this.SaveToFile("headimg",fileToName+fileName)
	var user models.Users
	o :=orm.NewOrm()
	user.Id =uid
	user.Avatar =fileToName+fileName
	o.Update(&user)
	result["code"]=200
	result["msg"]="修改成功！"
	this.Data["json"]=&result
	this.ServeJSON()
}
func authcode(str,operation,key string, expiry int64) string{		
	key =md5s(str)		
	keya :=md5s(substr(key,0,16))
	keyb := md5s(substr(key, 16, 16));
	var keyc string	 
	if operation =="DECODE"{
		keyc =substr(str,4,len(str))
	}else{
		keyc =substr(md5s(strconv.FormatInt(time.Now().UnixNano() / 1000000,10)),-4,0)
	}	
	cryptkey := keya + md5s(keya + keyc)
	
	key_length := len(cryptkey)	
	if operation =="DECODE"{
		stringss,_:=base64.StdEncoding.DecodeString(substr(str,4,len(str)))
		str =string(stringss)
	}else{
		if expiry !=0 {
			expiry =expiry+time.Now().Unix()
		}
		expirys :=fmt.Sprintf("%10d",expiry)
		expirys =expirys+substr(md5s(str+keyb),0,16)+str
	}
	string_length := len(str);
    var result string;
    box := make(map[int]int)
	rndkey :=make(map[int]string)
	for j:=0;j<=255;j++ {
		box[j]=j
	}
	for i:=0; i<=255; i++ {		
		rndkey[i]=string([]byte(cryptkey)[i%key_length])
	}
	for i:=0;i<256;i++{
		j:=i
		num,_:=strconv.Atoi(rndkey[i])
        j = (j + box[i] +num ) % 256
        tmp := box[i]
        box[i] = box[j]
        box[j] = tmp
	}
	for i := 0; i < string_length; i++ {
		a :=i
		j :=i
        a = (a + 1) % 256
        j = (j + box [a]) % 256
        tmp := box [a]
        box [a] = box [j]
		box [j] = tmp
		 num,_ :=strconv.Atoi(string([]byte(str)[i]))
        result = strconv.Itoa(num ^ box[(box[a] + box[j]) % 256])
	}
	if operation == "DECODE" {		
		return substr(result, 26,len(result));
	}else{
		stringss,_:=base64.StdEncoding.DecodeString(result)       
        return string(stringss)
	}
		
    
	
}
//md5加密
func md5s(key string) (keys string){	
	data := []byte(key)
    has := md5.Sum(data)
    md5str := fmt.Sprintf("%x", has)
	return md5str
}
//截取字符串
func substr(str string,start,end int)string{
	rs := []rune(str)
	if start <0{
		return string(rs[len(str)-4:len(str)])
	}
	return string(rs[start:end])
}