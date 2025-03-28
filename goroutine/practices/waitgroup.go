package practices

import (
	"fmt"
	"sync"
)

func sayHello(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Hello")
}

func N_170() {
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
