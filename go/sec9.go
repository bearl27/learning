package main

import (
	"fmt"
	"time"
)

func sec9() {
	fmt.Println("sec9")
	fmt.Println("--slice--")
	f_slice()
	fmt.Println("--map--")
	f_map()
	fmt.Println("--channel--")
	f_channel()
	// fmt.Println("--channel_go--")
	// f_channel_go()
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

func f_channel() {
	var ch1 chan int
	//　受信専用
	// var ch2 <-chan int
	// 送信専用
	// var ch3 chan<- int

	ch1 = make(chan int)
	ch2 := make(chan int)

	fmt.Println(cap(ch1))
	fmt.Println(cap(ch2))

	ch3 := make(chan int, 3)
	fmt.Println(cap(ch3)) // 3
	ch3 <- 1
	ch3 <- 2
	ch3 <- 3
	fmt.Println(len(ch3)) //3
	i := <-ch3
	fmt.Println("value", i, "len", len(ch3))
	fmt.Println("value", <-ch3, "len", len(ch3))
	ch3 <- 4
	ch3 <- 5
	// ch3 <- 6
	fmt.Println("value", <-ch3, "len", len(ch3))

	close(ch3)
	//ch3 <- 6 // panic
	i, ok := <-ch3
	fmt.Println(i, ok)

	ch4 := make(chan int, 2)
	close(ch4)
	i, ok = <-ch4
	fmt.Println(i, ok)

	ch5 := make(chan int, 2)
	ch6 := make(chan int, 2)

	ch6 <- 1

	// fmt.Println(<-ch5)
	// fmt.Println(<-ch6)

	select {
	case v1 := <-ch5:
		fmt.Println(v1)
	case v2 := <-ch6:
		fmt.Println(v2)
	default:
		fmt.Println("default")
	}

}

func f_channel_go() {
	ch_go1 := make(chan int)
	ch_go2 := make(chan int)

	//go reciever(ch_go1)
	//go reciever(ch_go2)

	i := 0
	for i < 10 {
		ch_go1 <- i
		ch_go2 <- i
		time.Sleep(50 * time.Millisecond)
		i++
	}

}

func reciever(c chan int) {
	for {
		i := <-c
		fmt.Println(i)
	}
}
