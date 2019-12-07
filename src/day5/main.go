package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type Message struct {
	Id   int
	Name string
}

func main() {
	message := make(chan Message, 100)
	result := make(chan error, 100)
	for i := 0; i < 3; i++ {
		go worker(i, message, result)
	}
	total := 0
	for k := 1; k <= 3; k++ {
		message <- Message{Id: k, Name: "job" + strconv.Itoa(k)}
		total += 1
	}
	close(message)
	for j := 1; j <= total; j++ {
		res := <-result
		if res != nil {
			fmt.Println(res.Error())
		}
	}
	close(result)
}
func worker(worker int, msg <-chan Message, result chan<- error) {
	for job := range msg {
		fmt.Println("worker:", worker, "msg:", job.Id, ":", job.Name)
		time.Sleep(time.Second * time.Duration(RandInt(1, 3)))
		result <- nil
	}

}
func RandInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Intn(max-min+1)
}
