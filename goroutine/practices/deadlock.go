package practices

import (
	"fmt"
	"sync"
	"time"
)

// deadlock

type value struct {
	mu    sync.Mutex
	value int
	name  string
}

func N_175() {
	var wg sync.WaitGroup
	printSum := func(v1, v2 *value) {
		defer wg.Done()
		v1.mu.Lock()
		fmt.Printf("v1 %v lock\n", v1.name)
		defer v1.mu.Unlock()

		time.Sleep(2 * time.Second)
		v2.mu.Lock()
		fmt.Printf("v2 %v lock\n", v2.name)
		defer v2.mu.Unlock()

		fmt.Println("sum:", v1.value+v2.value)
	}

	var a value = value{name: "a"}
	var b value = value{name: "b"}
	wg.Add(2)
	go printSum(&a, &b)
	go printSum(&b, &a)
	wg.Wait()
	fmt.Println("done")
}
