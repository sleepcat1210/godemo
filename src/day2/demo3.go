package main

import "fmt"

type Dog struct {
	Name string
}

func main() {
	var i interface{}
	var d = Dog{
		Name: "大狗狗",
	}
	i = d
	fmt.Println(i)
	i = "string"
	fmt.Println(i)
	i = false
	fmt.Println(i)
	i = 100
	fmt.Println(i)
	i = 3.1415
	fmt.Println(i)
}
