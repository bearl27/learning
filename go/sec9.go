package main

import "fmt"

func sec9() {
	fmt.Println("sec9")
	fmt.Println("--slice--")
	f_slice()
	fmt.Println("--map--")
	f_map()
}

func f_slice() {
	var sl []int
	fmt.Println(sl)
	fmt.Println(len(sl))

	var sl2 []string = []string{"a", "b"}
	fmt.Println(sl2)
	fmt.Println(len(sl2))

	sl3 := make([]int, 5)
	fmt.Println(sl3)
	fmt.Println(len(sl3))

	sl3[0] = 1
	fmt.Println(sl3)
	fmt.Println(sl3[2:4])

	// append
	sl2 = append(sl2, "c")
	fmt.Println(sl2)
	sl2 = append(sl2, "d", "e")
	fmt.Println(sl2)

	sl4 := make([]int, 2, 3)
	fmt.Println(len(sl4)) // 2
	fmt.Println(cap(sl4)) // 3
	sl4 = append(sl4, 1)
	fmt.Println(len(sl4)) // 3
	fmt.Println(cap(sl4)) // 3
	sl4 = append(sl4, 1)
	fmt.Println(len(sl4)) // 4
	fmt.Println(cap(sl4)) // 6 capを超える場合は倍増する
	sl4 = append(sl4, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(len(sl4)) // 14
	fmt.Println(cap(sl4)) // 14 前のcapの倍も超える場合は追加した要素数をふやす
	sl4 = append(sl4, 1)
	fmt.Println(len(sl4)) // 6
	fmt.Println(cap(sl4)) // 6

	// copy
	sl5 := []int{1, 2}
	sl6 := make([]int, 5, 10)
	copy_sl := copy(sl6, sl5)
	fmt.Println(sl6)
	fmt.Println(copy_sl)

	// for
	sl7 := []string{"a", "b", "c"}
	for k, v := range sl7 {
		fmt.Println(k, v)
	}
	for _, v := range sl7 {
		fmt.Print(v)
	}
	fmt.Println()
	for k := range sl7 {
		fmt.Print(k)
	}
	fmt.Println()

	fmt.Println(Sum(1, 2, 3))
	sl8 := []int{1, 2, 3}
	fmt.Println(Sum(sl8...))
	//fmt.Println(sl8...)
}

func Sum(s ...int) int {
	n := 0
	for _, v := range s {
		n += v
	}
	return n
}

func f_map() {
	var m map[int]string
	fmt.Println(m)
	var m1 = map[int]string{1: "a", 2: "b"}
	fmt.Println(m1)
	m2 := map[int]string{
		1: "a",
		2: "b",
	}
	fmt.Println(m2)
	fmt.Println(m2[1])
	fmt.Println(m2[3]) // 空白
	// 見分ける
	v, ok := m2[3]
	fmt.Println(v, ok)
	m2[8] = "c"
	fmt.Println(m2)
	delete(m2, 8)
	fmt.Println(len(m2))
	m3 := make(map[int]string)
	fmt.Println(m3)
}
