package main
import(
	"fmt"
	"strconv"
	"net/http"
	"os"
)
func HttpGet(url string)(result string,err error){
	resp,err1 :=http.Get(url)
	if err1 !=nil{
		err =err1;
		return
	}
	defer resp.Body.Close()
	buf :=make([]byte,4096)
	for{
		n,err :=resp.Body.Read(buf)
		if n ==0{
			fmt.Println("读完！err:",err)
			break
		}
		result +=string(buf[:n])
	}
	return

}
func working(start,end int){
	fmt.Printf("正在爬取%d到%d页\n",start,end)
	for i:=start;i<=end;i++{
		url :="https://tieba.baidu.com/f?kw=%E7%BB%9D%E5%9C%B0%E6%B1%82%E7%94%9F&ie=utf-8&pn="+
		strconv.Itoa((i-1)*50)
		fmt.Printf("正在爬：%d 页,%s\n",i,url)
		result,err :=HttpGet(url)
		if err !=nil{
			fmt.Println("httpGet err:",err)
			continue
		}
		fileName :=strconv.Itoa(i)+".html"
		f,err :=os.Create(fileName)
		if err !=nil{
			fmt.Println("create err:",err)
			continue
		}
		f.WriteString(result)
		f.Close()
	}
}

func main() {
	//指定爬取的起始位置，终止位置
	var start,end int
	fmt.Println("请输入爬取的起始页面：")
	fmt.Scan(&start)
	fmt.Println("请输入爬取终止页面：")
	fmt.Scan(&end)
	working(start,end)
    
}
