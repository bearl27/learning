package main

import "fmt"

// HelloWorld
// func main() {
// 	fmt.Println("Hello World")
// 	fmt.Println(time.Now())
// }

// 変数
func main() {
	//func_var()
	//func_int()
	//func_float()
	//func_string()
	// func_interface()
	//func_const()
	//sec7()
	//sec8()
	//sec9()
	sec10()
}

func func_var() {
	// 明示的な型定義
	var i int = 100
	fmt.Println(i)
	var s string = "Hello"
	fmt.Println(s)

	var (
		i2 int    = 200
		s2 string = "Go"
	)
	fmt.Println(i2, s2)

	// 暗黙的な型定義
	i3 := 300
	fmt.Println(i3)
	s3 := "Go2"
	fmt.Println(s3)

	// 型のみの宣言
	var i4 int
	var s4 string
	fmt.Println(i4, s4)
	i4 = 400
	s4 = "Go3"
	fmt.Println(i4, s4)

	outer()
	// s4はGo3のまま
	fmt.Println(s4)
}

func outer() {
	var s4 string = "outer"
	fmt.Println(s4)
}

func func_int() {
	var i int = 100
	fmt.Println(i)
	var i2 int64 = 200
	fmt.Println(i2)
	// error
	// fmt.Println(i + i2)
	// 変換
	fmt.Println(i + int(i2))
	fmt.Printf("%T\n", i)
	fmt.Printf("%T\n", i2)
}

func func_float() {
	var fl64 float64 = 2.4
	fmt.Println(fl64)
	fl := 3.2
	fmt.Println(fl)

	var fl32 float32 = 1.2
	fmt.Println(fl32)

	fmt.Printf("%T\n", fl64)
	fmt.Printf("%T\n", fl)
	fmt.Printf("%T\n", fl32)

	zero := 0.0
	pinf := 1.0 / zero
	minf := -1.0 / zero
	nan := zero / zero
	fmt.Println(pinf)
	fmt.Println(minf)
	fmt.Println(nan)
}

func func_string() {
	var s string = "Hello Golang"
	fmt.Println(s)
	fmt.Printf("%T\n", s)
	var sn = `test
	test
		test`
	fmt.Println(sn)
	fmt.Println(s[0])
	fmt.Println(string(s[0]))
}

func func_interface() {
	var x interface{}
	fmt.Println(x)
	x = 1
	fmt.Println(x)
	var y interface{}
	y = 2
	fmt.Println(y)
	//fmt.Println(x + y)
}

var num = 1
var (
	i    int = 1
	f64      = 1.2
	s        = "test"
	t, f     = true, false
)

const (
	c0 = 0
	c1
	c2
	c3
)

const (
	c00 = iota
	c01
	c02
	c03
)

const Big = 9999999999999999999999

func func_const() {
	fmt.Println(num)
	fmt.Printf("%T\n", num)
	fmt.Println(i, f64, s, t, f)
	fmt.Println(c0, c1, c2, c3)
	fmt.Println(c00, c01, c02, c03)
	//fmt.Println(Big)
	fmt.Println(Big - 10000000000000000000000)
}
