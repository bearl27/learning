package main

import "fmt"

func Plus(a int, b int) int {
	return a + b
}

func Div(x, y int) (int, int) {
	q := x / y
	r := x % y
	return q, r
}

func Double(price int) (result int) {
	result = price * 2
	return
}

func NoReturn() {
	fmt.Println("NoReturn")
}

func add(x int) {
	x++
}

func sec7() {
	fmt.Println("sec7")
	i := Plus(1, 2)
	fmt.Println(i)

	i2, i3 := Div(5, 2)
	fmt.Println(i2, i3)

	i2, _ = Div(3, 2)
	fmt.Println(i2)

	i4 := Double(2)
	fmt.Println(i4)

	NoReturn()

	x := 1
	add(x)
	fmt.Println(x)

	// noname func
	f := func(x, y int) int {
		return x + y
	}
	fmt.Println(f(1, 2))

	f1 := func(x, y int) int {
		return x - y
	}(1, 2)
	fmt.Println(f1)

	f3 := returnFunc()
	f3()
	callFunc(func() {
		fmt.Println("callFunc")
	})

	f4 := Later()
	fmt.Println(f4("Hello"))
	fmt.Println(f4("World"))
	fmt.Println(f4("!"))
	ints := integers()
	fmt.Println(ints())
	fmt.Println(ints())
	otherInts := integers()
	fmt.Println(otherInts())
	fmt.Println(otherInts())
}

// return func
func returnFunc() func() {
	return func() {
		fmt.Println("returnFunc")
	}
}

// call func
func callFunc(f func()) {
	f()
}

// クロージャー（CLOJURE）
func Later() func(string) string {
	var store string
	return func(next string) string {
		s := store
		store = next
		return s
	}
}

// generater
func integers() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
