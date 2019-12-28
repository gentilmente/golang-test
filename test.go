package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func say(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(time.Millisecond * 30)
	}
	wg.Done()
}

func main() {
	wg.Add(1)
	go say("hello")
	wg.Add(1)
	go say("goodbye")
	wg.Wait()

	wg.Add(1)
	go say("hi")
	wg.Wait()
}
