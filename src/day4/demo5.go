package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

//互斥枷锁
func main() {
	var state = make(map[int]int)
	var mutex = &sync.Mutex{}
	var ops int64 = 0
	for r := 0; r < 100; r++ {
		total := 0
		go func() {
			key := rand.Intn(5)
			mutex.Lock()
			total += state[key]
			mutex.Unlock()
			atomic.AddInt64(&ops, 1)
			runtime.Gosched()
		}()
	}
	//10 协程
	for w := 0; w < 10; w++ {
		go func() {
			for {
				key := rand.Intn(5)
				val := rand.Intn(100)
				mutex.Lock()
				state[key] = val
				mutex.Unlock()
				atomic.AddInt64(&ops, 1) //在ops 增加1
				runtime.Gosched()
			}
		}()
	}
	time.Sleep(time.Second)
	opsFinal := atomic.LoadInt64(&ops)
	fmt.Println("ops:", opsFinal)
	mutex.Lock()
	fmt.Println("state:", state)
	mutex.Unlock()
	fmt.Println(rand.Intn(5)) //随机int 0-n
}
