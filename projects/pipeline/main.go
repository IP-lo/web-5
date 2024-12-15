package main

import "fmt"

func removeDuplicates(inputStream chan string, outputStream chan string) {
	defer close(outputStream)
	var prev string

	for value := range inputStream {
		if value != prev {
			outputStream <- value
			prev = value
		}
	}
}

func main() {
	var strIn string 
	fmt.Scan(&strIn)
	in := make(chan string)
	out := make(chan string)
	go removeDuplicates(in, out)
	go func ()  {
		for _, val:= range strIn{
			in <- string(val)
		}
		close(in)
	}()

	for val := range out{
		fmt.Print(val)
	}

}
