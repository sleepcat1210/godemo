package main

import (
	"fmt"
	_ "fmt"
	"io/ioutil"
	"log"
	_ "log"
	"strings"
)

func main() {
	// file,err:=os.Open("aa.txt")
	// if err !=nil{
	// 	log.Fatal(err)
	// }
	// data :=make([]byte,100)
	// count,err :=file.Read(data)
	// if err !=nil{
	// 	log.Fatal(err)
	// }
	// fmt.Printf("read %d bytes:%q\n",count,data[:count])
	// file.Close()
	// file,_ :=os.Stat("aa.txt")
	// fmt.Println(file.Name(),file.IsDir(),file.ModTime())
	// os.Mkdir(`dir`,0777)
	// os.MkdirAll(`a/b/c/d`, 0777)
	// os.Rename(`dir`, `dir1`) //修改一个文件或者移动文件
	// os.Remove(`dir1`) //删除
	// os.RemoveAll(`a`) //删除多个文件夹
	// f1, _ := os.Create(`demo.go`)
	// f2, _ := os.Open(`aa.txt`) //只读
	// io.Copy(f1, f2)            //io数据流
	//
	// font, _ := os.OpenFile(`output.txt`, os.O_RDWR, 0777) //可以对文件读写
	// defer font.Close()
	// font.WriteString(`的建安街道11212`)
	// r := strings.NewReader("gi is  dsd f dsfs fs f sf sd fs f s fds f")
	// b, err := ioutil.ReadAll(r)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%s", b)
	// f, err := ioutil.ReadFile(`demo.go`) //读取文件的数据
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%s", f)
	// ioutil.WriteFile(`a.txt`, []byte("上海"), 0777) //所写入文件 byte 权限模式
	// files, err := ioutil.ReadDir(".")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// for _, file := range files {
	// 	if file.IsDir() {
	// 		file1, _ := ioutil.ReadDir("./" + file.Name())
	// 		for _, f := range file1 {
	// 			fmt.Println(f.Name())
	// 		}

	// 	}
	// 	fmt.Println(file.Name())
	// }
	// readdir("./src/day1")
	// fmt.Println(strings.Compare("1", "-1")) //1
	// fmt.Println(strings.Compare("-1", "1"))//-1
	// fmt.Println(strings.Compare("1", "1"))//0
	// fmt.Println(strings.Contains("aabbddsdad", "ds")) //是否包含
	// fmt.Println(strings.Count("abcdsdsds", "s")) //返回包含的个数
	// fmt.Println(strings.Index("aaacdsfsfsfsde", "s")) //返回出现位置 5
	// s := []string{"zhangsan", "lisi", "wangpo"}
	// fmt.Println(strings.Join(s, ","))
	// fmt.Println(strings.Replace("s a b c c c", "c", "a", 2)) //替换字符串的指定的字符 并指出替换多少个
	// fmt.Printf("%q\n", strings.Split("s,d,s,f", ",")) //分割字符串,
	fmt.Println(strings.ToUpper("g")) //转化大写
	fmt.Println(strings.ToLower("G")) //转化小写
}
func readdir(dir string) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		if file.IsDir() {
			readdir(dir + "/" + file.Name())
		}

		fmt.Println(file.Name() + "\n")

	}
}
