package practices

import (
	"fmt"
	"sync"
	"time"
)

func N_179() {
	var mutex sync.Mutex
	cond := sync.NewCond(&mutex)

	for _, name := range []string{"A", "B", "C"} {
		go func(name string) {
			mutex.Lock()
			defer mutex.Unlock()
			cond.Wait()
			fmt.Println(name)
		}(name)
	}

	fmt.Println("Ready..")
	time.Sleep(time.Second)
	fmt.Println("Go")
	// for i := 0; i < 3; i++ {
	// 	time.Sleep(time.Second)
	// 	cond.Signal()
	// }

	cond.Broadcast()

	time.Sleep(time.Second)
	fmt.Println("Done")
}
