package main

import (
	"fmt"
	"strconv"
)

func PrintSliceInts(i []int) {
	for _, v := range i {
		fmt.Println(v)
	}
}

func PrintSliceStrings(i []string) {
	for _, v := range i {
		fmt.Println(v)
	}
}

func PrintSlice[T any](s []T) {
	for _, v := range s {
		fmt.Println(v)
	}
}

func f[T fmt.Stringer](xs []T) []string {
	result := []string{}
	for _, x := range xs {
		result = append(result, x.String())
	}
	return result
}

type Myint int

type Vector[T any] []T

func (i Myint) String() string {
	return strconv.Itoa(int(i))
}

type Number interface {
	~int | ~int32 | ~int64 | ~float32 | ~float64
}

func Max[T Number](a, b T) T {
	if a >= b {
		return a
	}
	return b
}

type T[A any, B []C, C *A] struct {
	a A
	b B
	c C
}

type Set[T comparable] map[T]struct {
}

func NewSet[T comparable](items ...T) Set[T] {
	s := make(Set[T])
	for _, item := range items {
		s[item] = struct{}{}
	}
	return s
}

func main() {
	// PrintSliceInts([]int{1, 2, 3, 4, 5, 6})
	// PrintSliceStrings([]string{"a", "b", "c", "d", "e", "f"})
	// PrintSlice[int]([]int{1, 2, 3, 4, 5, 6})
	// PrintSlice[string]([]string{"a", "b", "c", "d", "e", "f"})
	// fmt.Println(f([]Myint{1, 2, 3, 4, 5, 6}))
	// fmt.Println(Max(1, 2))
	// fmt.Println(Max(1.1, 2.2))
	// var x, y Myint = 1, 2
	// fmt.Println(Max(x, y))

	// var v Vector[int] = []int{1, 2, 3, 4, 5, 6}
	// fmt.Println(v)

	// var t T[int, []*int, *int]
	// fmt.Printf("A: %T, B: %T, C: %T", t.a, t.b, t.c)
}
