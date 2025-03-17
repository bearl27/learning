package practices

import (
	"fmt"
	"sync"
	"time"
)

func N_178() {
	var count int
	var lock sync.RWMutex
	var wg sync.WaitGroup

	increment := func(wg *sync.WaitGroup, l sync.Locker) {
		l.Lock()
		defer l.Unlock()
		defer wg.Done()

		fmt.Println("increment")
		count++
		time.Sleep(time.Second)
	}

	read := func(wg *sync.WaitGroup, l sync.Locker) {
		l.Lock()
		defer l.Unlock()
		defer wg.Done()

		fmt.Println("read")
		fmt.Println(count)
		time.Sleep(time.Second)
	}

	start := time.Now()
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go increment(&wg, &lock)
	}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go read(&wg, lock.RLocker())
	}
	wg.Wait()
	// end := time.Now()
	// fmt.Println(end.Sub(start))
	fmt.Println(time.Since(start))
}
