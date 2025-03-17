package practices

import (
	"fmt"
	"sync"
	"time"
)

func N_173() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("1st goroutine")
		time.Sleep(1 * time.Second)
		fmt.Println("1st goroutine finished")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("2nd goroutine")
		time.Sleep(1 * time.Second)
		fmt.Println("2nd goroutine finished")
	}()

	wg.Wait()
}
