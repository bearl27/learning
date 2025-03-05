package main

import (
	"fmt"
	"strconv"
)

func sec8() {
	println("sec8")
	f_if()
}

func f_if() {
	a := 0
	if a == 0 {
		println("a is zero")
	} else if a == 1 {
		println("a is one")
	} else {
		println("a is not zero, one")
	}

	if b := 100; b == 100 {
		println("b is 100")
	}

	x := 0
	if x := 2; true {
		println(x)
	}
	println(x)

	var s string = "A"

	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("i = %d\n", i)
}
