package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func dowork(d time.Duration, resch chan string) {
	fmt.Println("Doing work...")
	time.Sleep(d)
	fmt.Println("Work Done")
	resch <- fmt.Sprintf("the result of the work -> %d", rand.Intn(100))
	wg.Done()
}

var wg *sync.WaitGroup

func main() {
	start := time.Now()
	resultch := make(chan string)
	wg = &sync.WaitGroup{}
	wg.Add(3)
	go dowork(time.Second*2, resultch)
	go dowork(time.Second*4, resultch)
	go dowork(time.Second*6, resultch)

	go func() {
		for res := range resultch {
			fmt.Println(res)
		}
		fmt.Printf("work took %v seconds\n", time.Since(start))
	}()

	wg.Wait()
	close(resultch)
	time.Sleep(time.Second)
}
