package practices

import (
	"fmt"
	"sync"
)

var once sync.Once

func A() {
	fmt.Println("A")
	once.Do(B)
}
func B() {
	fmt.Println("A")
}

func N_181() {
	once.Do(A)
}
