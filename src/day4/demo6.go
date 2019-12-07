package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)
//枷锁
func main() {
	var ops int64 = 0
	var total int
	var mutex = &sync.Mutex{}
	for i := 1; i < 100; i++ {
		go func() {
			mutex.Lock()
			total += 1
			mutex.Unlock()
		}()

	}
	for i := 1; i < 100; i++ {
		go func() {
			mutex.Lock()
			total += 1
			mutex.Unlock()
		}()

	}
	atomic.AddInt64(&ops, 1)
	opsFinal := atomic.LoadInt64(&ops)
	fmt.Println(opsFinal)
	fmt.Println(total)
}
