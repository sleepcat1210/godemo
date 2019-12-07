package main
import(
	"fmt"
	"os"
	"net"
	"io"
)
func sendFile(conn net.Conn,path string){
	//只读方式打开文件
	f,err :=os.Open(path)
	if err !=nil{
		fmt.Println("os.Open err:",err)
		return
	}
	defer f.Close()
	buf :=make([]byte,4096)
	for {
		n,err :=f.Read(buf)
		if err !=nil{
			if err == io.EOF{
				fmt.Println("文件发送完毕")
			}else{
				fmt.Println("f.Read err:",err)
			}
			return
		}
		conn.Write(buf[:n])
	}
}


func main() {
	fmt.Println("请输入路径")
	var path string
	fmt.Scan(&path)
	//获取文件名
	fileInfo,err :=os.Stat(path)
	if err !=nil{
		fmt.Println("os stat err:",err)
		return
	}
	//主动连接服务器
	conn,err :=net.Dial("tcp","127.0.0.1:9090")
	if err !=nil{
		fmt.Println("dial err:",err)
		return
	}
	defer conn.Close()
	//给接收端 发送文件名、
	_,err =conn.Write([]byte(fileInfo.Name()))
	if err !=nil{
		fmt.Println("conn write err :",err)
		return
	}
	//读取接收端回发
	buf :=make([]byte,1024)
	n,err :=conn.Read(buf)
	if err !=nil{
		fmt.Println("conn.Read err:",err)
		return
	}
	if "ok" ==string(buf[:n]){
		sendFile(conn,path)
	}
}