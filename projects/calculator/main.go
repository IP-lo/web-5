package main

import (
	"fmt"
)
func calculator(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int {
    resChan := make(chan int)
    go func(){
		defer close(resChan)
		select{
		case val, _ := <-firstChan:
			resChan <- val * val
		case val, _ := <-secondChan:
			resChan <- val * 3
		case <-stopChan:
			return
		}
	}()
	return resChan
}

func main() {
	firstChan := make(chan int)
	secondChan := make(chan int)
	stopChan := make(chan struct{})
	resChan := calculator(firstChan, secondChan, stopChan)

	go func() {
		firstChan <-4
		close(firstChan)
	}() 
	go func() {
		secondChan <-5
		close(secondChan)
		close(stopChan)
	}() 

	for res := range resChan{
		fmt.Println(res)
	}

}
