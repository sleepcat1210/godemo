package main

import "fmt"

// func main() {
// 	defer func() {
// 		if err := recover(); err != nil {
// 			fmt.Println("defer:", err)
// 		}
// 	}()
// 	panic("提前终止程序")
// }
func fn1() {
	panic("painc")
}
func fn2() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("defer:", err)
		}
	}()
	fmt.Println("zzzadadadfe")
}
func main() {
	for i := 0; i < 3; i++ {
		fn1()
		fn2()
	}
}
