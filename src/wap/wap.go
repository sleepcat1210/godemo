package main
import(
	"fmt"
	"net/http"
	"log"
	"os"
	"io"
	"strconv"
	
)
func  say(w http.ResponseWriter, r *http.Request){
	if r.Method =="POST"{
		f,h,err := r.FormFile("file")
		if err !=nil{
			w.Write([]byte("上传文件错误："+err.Error()))
			return
		}
		
		os.Mkdir("./static",0777)//创建目录
		out,err:=os.Create("./static/"+h.Filename)//创建文件返回句柄
		if err !=nil{
			io.WriteString(w,"文件创建失败："+err.Error())
			return
		}
		_,err =io.Copy(out,f)
		if err !=nil{
			io.WriteString(w,"文件保存失败："+err.Error())
			return
		}
		io.WriteString(w,"/static/"+h.Filename)
		return
	}
	if r.Method=="GET"{
		io.WriteString(w,"测试get请求")
		return
	}
  
}
//io 输出输入流  os 文件操作
func imageview(w http.ResponseWriter, r *http.Request){
	r.ParseForm()//存放到容器中
	name:=r.Form.Get("name")//获取name参数
	f,err :=os.Open("./static/"+name)//打开文件获取句柄
	if err !=nil{
		w.WriteHeader(404)
		return
	}
	defer f.Close()//关闭
	w.Header().Set("Content-Type","image")//设置输出格式
	io.Copy(w,f)
}
//获取文件信息
func getInfo(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	pathName :=r.Form.Get("name")
	f,err :=os.Stat("./static/"+pathName)
	if err !=nil{
		fmt.Println("文件错误：",err.Error())
		return
	}
	name :=f.Name()
	size :=strconv.FormatInt(f.Size(),10)
	io.WriteString(w,name)
	io.WriteString(w,size)
	path,_ :=os.Getwd()//获取文件位置
	io.WriteString(w,path)
	fmt.Println(size)
	fmt.Println(size)
}
func main() {
	http.HandleFunc("/info",getInfo)
	http.HandleFunc("/",say)		
	err :=http.ListenAndServe(":9090",nil)
	if err !=nil{
		log.Fatal("listenAndServer: ",err)
	}
}