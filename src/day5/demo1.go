package main

import (
	"fmt"
	"time"
)

func main() {
	time1 := time.NewTicker(time.Second * 2)
	<-time1.C
	fmt.Println("Timer 1 expired")
	timer2 := time.NewTicker(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("timer 2 expired")
	}()
	// stop2 := timer2.Stop()
	// if timer2.Stop() {
	// 	fmt.Println("timer 2 stopped")
	// }
}
