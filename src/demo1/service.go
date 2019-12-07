package main
import(
	"fmt"
	"net"
	"bufio"
)
func process(conn net.Conn){
	defer conn.Close()
	//建立缓冲区
	for{
		read :=bufio.NewReader(conn)//获取read指针
		var buf [1024]byte
		n,err :=read.Read(buf[:])
		if err != nil{
			fmt.Println("read err ",err)
			break
		}
		str :=string(buf[:n])
		conn.Write([]byte(str))
	}
	
}
func main(){
	//建立监听
	listen,err :=net.Listen("tcp","127.0.0.1:8000")
	if err !=nil{
		fmt.Println("listen err ",err)
		return
	}
	for {
		conn,err :=listen.Accept()
		if err != nil{
			fmt.Println("accept err ",err)
			continue
		}
		go process(conn)
	}
}