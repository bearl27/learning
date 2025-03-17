package practices

import (
	"sync"
	"time"
)

func N_180() {
	cond := sync.NewCond(&sync.Mutex{})

	go func() {
		for range time.Tick(time.Second) {
			cond.Broadcast()
		}
	}()

	var flag [2]bool

	takeStop := func() {
		cond.L.Lock()
		defer cond.L.Unlock()
		cond.Wait()
	}

	var wg sync.WaitGroup

	p0 := func() {
		defer wg.Done()
		flag[0] = true
		takeStop()
		for flag[1] {
			takeStop()

			flag[0] = false
			takeStop()
			if flag[0] != flag[1] {
				break
			}
			takeStop()

			flag[0] = true
			takeStop()
		}
	}
	p1 := func() {
		defer wg.Done()
		flag[1] = true
		takeStop()
		for flag[0] {
			takeStop()

			flag[1] = false
			takeStop()
			if flag[0] != flag[1] {
				break
			}
			takeStop()

			flag[1] = true
			takeStop()
		}
	}

	wg.Add(2)
	go p0()
	go p1()
	wg.Wait()
}
