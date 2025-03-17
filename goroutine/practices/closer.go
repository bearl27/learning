package practices

import (
	"fmt"
	"sync"
)

func N_171() {
	// var wg sync.WaitGroup
	// say := "Hello"
	// wg.Add(1)

	// go func() {
	// 	defer wg.Done()
	// 	say = "Good Bye"
	// }()

	// wg.Wait()
	// fmt.Println(say)
	pre()
}

func pre() {
	var wg sync.WaitGroup

	tasks := []string{"A", "B", "C"}

	for _, task := range tasks {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(task)
		}()
	}
	wg.Wait()
}
