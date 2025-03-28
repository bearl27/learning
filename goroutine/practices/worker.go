package practices

import (
	"fmt"
	"sync"
	"time"
)

func N_177() {
	var wg sync.WaitGroup
	var lock sync.Mutex

	const timer = 1 * time.Second

	greedyWorker := func() {
		defer wg.Done()

		count := 0
		begin := time.Now()
		for time.Since(begin) <= timer {
			lock.Lock()
			time.Sleep(3 * time.Nanosecond)
			lock.Unlock()

			count++
		}

		fmt.Println("greed", count)
	}

	politeWorker := func() {
		defer wg.Done()

		count := 0
		begin := time.Now()
		for time.Since(begin) <= timer {
			lock.Lock()
			time.Sleep(1 * time.Nanosecond)
			lock.Unlock()
			lock.Lock()
			time.Sleep(1 * time.Nanosecond)
			lock.Unlock()
			lock.Lock()
			time.Sleep(1 * time.Nanosecond)
			lock.Unlock()

			count++
		}
		fmt.Println("polite:", count)
	}

	wg.Add(2)
	go greedyWorker()
	go politeWorker()
	wg.Wait()
}
