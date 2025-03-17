package main

import (
	"fmt"
	"sync"
	"time"
)

func sayHello(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Hello")
	time.Sleep(1 * time.Second)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	go sayHello(&wg)
	wg.Add(1)

	go func() {
		defer wg.Done()
		fmt.Println("Bye")
	}()
	fmt.Println("waiting")
	wg.Wait()
	fmt.Println("End")
}
