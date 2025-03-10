package main

import "fmt"

func Double(i int) {
	i = i * 2
}

func DoublePointer(i *int) {
	*i = *i * 2
}

// pointa
func sec10() {
	fmt.Println("--Pointa--")
	n := 100
	fmt.Println(n)
	fmt.Println(&n) // address
	Double(n)
	fmt.Println(n)
	var p *int = &n
	fmt.Println(p)
	fmt.Println(*p)

	*p = 300
	fmt.Println(n)
	n = 200
	fmt.Println(p)
	fmt.Println(*p)

	DoublePointer(p)
	fmt.Println(n)
}
