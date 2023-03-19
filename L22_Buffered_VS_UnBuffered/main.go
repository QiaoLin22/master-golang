package main

import "fmt"

func main() {
	msgch := make(chan int, 1000)

	go func() {
		msg := <-msgch
		fmt.Println(msg)
	}()

	msgch <- 10

}
