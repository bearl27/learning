package test

import "fmt"

func IsOne(n int) bool {
	return n == 1
}

func index() {
	num := 10
	fmt.Println(IsOne(num))
}
