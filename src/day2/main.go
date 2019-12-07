package main

import (
	"fmt"
	"sync"
	"time"
)

func fn1() {
	time.Sleep(1 * time.Second)
	fmt.Println("暂停 1s")
	//执行完成就关闭一个等待
	wg.Done()
}
func fn2() {
	time.Sleep(2 * time.Second)
	fmt.Println("暂停 2s")
	wg.Done()
}

//waitGroup 可用来等待写成执行完成
var wg sync.WaitGroup

func main() {
	//开始时间
	begin := time.Now()
	for i := 0; i < 100; i++ {
		go fn1()
		wg.Add(1)
		go fn2()
		wg.Add(1)
	}
	//等待所有子协程完成
	wg.Wait()
	//获取运行结束时间
	end := time.Now()
	fmt.Println("总共用时：", end.Sub(begin))
}
