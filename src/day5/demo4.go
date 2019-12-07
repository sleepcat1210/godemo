package main
import(
	"os"
	"fmt"
	"math/rand"
	"time"
)
func main() {
	os.MkdirAll("./img/a",0777)
	str :="sfjsflajflajflkafucia"
	types :=[]byte(str)
	fmt.Println(types[0])
	fmt.Println(string(types[0]))
	fmt.Println(len(types))
	fmt.Println(string(types))
	fmt.Println(gets("dd"))
	fmt.Println(GetString(20))
}
func gets(str string) string{
	return str
}
func  GetString(l int) string{
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	// bytes := []byte(str)
	// result := []byte{}
	// r := rand.New(rand.NewSource(time.Now().UnixNano()))
	// for i := 0; i < l; i++ {
	// 	result = append(result, bytes[r.Intn(len(bytes))])
	// }
	bytes :=[]byte(str)
	result :=[]byte{}
	r :=rand.New(rand.NewSource(time.Now().UnixNano()))
	for i:=0;i<l;i++{
		result =append(result,bytes[r.Intn(len(bytes))])//r.Intn()随机数
	}
	return string(result)

}