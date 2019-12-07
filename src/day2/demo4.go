package main

import (
	"fmt"
	"net/http"
)

type Point struct {
	X, Y float64
}

func (p Point) Format() {
	fmt.Println("x:", p.X, "y:", p.Y)
}
func main() {
	// var p Point
	// p.X = 1000
	// // p.Format()
	// p.Format()
	resp, err := http.Get(`https://blog.zxysilent.com/post/goweb-03-3.html`)
	if err != nil {
		panic(err)
	}
	//准备容器
	buf := make([]byte, 1024*100)
	l, err := resp.Body.Read(buf)
	fmt.Println(l, err)
	fmt.Println(string(buf[:l]))
}
