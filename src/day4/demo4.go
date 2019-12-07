package main

import "fmt"

type Callback func(x, y int) int

func addIn() (int, int) {
	return 1, 3
}
func test(x, y int, callback Callback) int {
	return callback(x, y)
}
func add(x, y int) int {
	return x + y
}
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
func main() {
	a, b := addIn()
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println("-------------------------------------")
	x, y := 1, 3
	fmt.Println(test(x, y, add))
	fmt.Println("----------------------------")
	fmt.Println(split(20))
}
