package main

import (
	"fmt"
	"time"
	"sync"
)

func work() {
	time.Sleep(time.Millisecond * 50)
	fmt.Println("done")
}

func main() {
	wg := new(sync.WaitGroup)
	for i := 0; i < 10; i++{
		wg.Add(1)
		go func(){
			defer wg.Done()
			work()
		}()
	}
	wg.Wait()
}
