package main
import(
	"net"
	"fmt"
	"bufio"
	"os"
	"strings"
)
func main() {
	//键连接
	conn,err :=net.Dial("tcp","127.0.0.1:8000")
	if err !=nil{
		fmt.Println("err",err)
		return
	}
	//关闭连接
	defer conn.Close()
	read :=bufio.NewReader(os.Stdin)
	for {
		input,_:=read.ReadString('\n')//一行一行读
		//除去换行符
		inputInfo :=strings.Trim(input,"\r\n")
		if strings.ToUpper(inputInfo)=="Q"{
			return
		}
		_,err :=conn.Write([]byte(inputInfo))
		if err !=nil{
			return
		}
		var buf [1024]byte
		n,err :=conn.Read(buf[:])
		if err !=nil{
			fmt.Println("err",err)
			return
		}
		fmt.Println(string(buf[:n]))
	}
	
}