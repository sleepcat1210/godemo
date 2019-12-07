package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Person struct {
	Name     string //
	Age      int
	Addr     string
	Birthday time.Time
}
type Point struct {
	X int
	Y int `json:"y"`
}
type Circle struct {
	Point      //圆心点
	Radius int //半径
}

func main() {
	args := map[string]int{
		"zxy": 24,
		"zzz": 24,
	}
	args["ceshi"] = 60
	args["zxy"]++
	fmt.Println(args)
	delete(args, "ceshi")
	fmt.Println(args)
	for k, v := range args {
		fmt.Println(k, v)
	}
	fmt.Println("--------------------------------------")
	for k := range args {
		fmt.Println(k)
	}
	fmt.Println("---------------------------------")
	for _, v := range args {
		fmt.Println(v)
	}
	fmt.Println("-------------------分割线------------------------")
	var zxy Person
	zxy.Name = "张三"
	fmt.Println(zxy)
	age := &zxy.Age
	*age = 8
	fmt.Println(zxy)
	lisi := &zxy
	lisi.Name = "李四"
	fmt.Println(zxy)
	fmt.Println("---------------------分割线---------------------------")
	c := Circle{
		Point:  Point{1, 0},
		Radius: 10,
	}
	c1 := Circle{
		Point{2, 0},
		10,
	}
	fmt.Println(c1.X)
	fmt.Println(c.Point.X)
	buf, _ := json.Marshal(c) //返回byte
	fmt.Println(string(buf))
	s := `{"X":1,"Y":2,"半径":10}`
	de := Circle{}
	json.Unmarshal([]byte(s), &de)
	fmt.Println(de)
}
