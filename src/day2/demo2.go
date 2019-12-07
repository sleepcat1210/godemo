package main

import (
	"fmt"
)

//定一个animal接口
type Animal interface {
	Say()
}

//具体实现类型去实现所定义的方法
type Dog struct {
	Name string
}

//say  实现所需的方法
func (d Dog) Say() {
	fmt.Println("Dog say :", d.Name)
}

//结构体
type Cat struct {
	Name string
}

func (c Cat) Say() {
	fmt.Println("Cat say:", c.Name)
}
func (c Cat) String() string {
	return fmt.Sprintf("\nstringer cat:%s", c.Name)
}
func main() {
	//一个接口变量
	var i Animal
	var d = Dog{
		Name: "大狗狗",
	}
	//实现了接口就可以把赋值给接口
	i = d
	d.Say()
	fmt.Printf("%T\n", i)
	fmt.Println(i)
	fmt.Println("------------------华丽分割线----------------------")
	//定义cat变量
	var c = Cat{
		Name: "小猫猫",
	}
	i = c
	c.Say()
	fmt.Printf("%T", i)
	fmt.Println(i)
}
