package sections

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func Sec8() {
	println("sec8")
	println("--if--")
	f_if()
	println("--for--")
	f_for()
	println("--switch--")
	f_switch()
	println("--interface--")
	f_interface()
	println("--label for--")
	f_label_for()
	println("--defar--")
	f_defar()
	println("--panic--")
	f_panic()
	println("--goroutine--")
	f_goroutine()

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

func f_for() {
	for i := 0; i < 10; i++ {
		fmt.Print(i)
	}
	fmt.Println()

	mario := "mario"
	count := 0
	for {
		count++
		if mario[count] == 'a' {
			fmt.Println(count)
			break
		}
		if count > len(mario) {
			fmt.Println("not found")
			break
		}
	}

	while := 0
	for while < 10 {
		fmt.Print(while)
		while++
	}
	fmt.Println()

	arr := [3]int{1, 2, 3}
	for i, v := range arr {
		fmt.Println(i, v)
	}

	sl := []string{"Python", "PHP", "GO"}
	fmt.Println(len(sl))
	for i, v := range sl {
		fmt.Println(i, v)
	}

	m := map[string]int{"Python": 1, "PHP": 2, "GO": 3}
	for k, v := range m {
		fmt.Println(k, v)
	}
}

func f_switch() {
	a := 0
	switch a {
	case 0, 2:
		println("a is zero or two")
	case 1:
		println("a is one")
	default:
		println("i dont no")
	}

	switch n := 2; n {
	case 1:
		println("a is one")
	case 2:
		println("a is two")
	default:
		println("i dont no")
	}

	n := 6
	switch {
	case n%2 == 0:
		println("n is even")
	case n%2 == 1:
		println("n is odd")
	}
}

func f_interface() {
	anything(1)
	anything("test")
	anything(true)

	var x interface{} = 3
	i, ok := x.(int)
	fmt.Println(i, ok)
	fmt.Println(i + 2)
	//fmt.Println(x + 2)

	//run time error
	// f := x.(float64)
	// fmt.Println(f)
	//　↓
	f2, ok2 := x.(float64)
	fmt.Println(f2, ok2)

	x = "test"

	if x == nil {
		fmt.Println("none")
	} else if i, ok := x.(int); ok {
		fmt.Println(i)
	} else if s, ok := x.(string); ok {
		fmt.Println(s, "string")
	} else {
		fmt.Println("i dont know")
	}

	switch x.(type) {
	case int:
		fmt.Println(x, "int")
	case string:
		fmt.Println(x, "string")
	default:
		fmt.Println(x, "i dont know")
	}

	switch v := x.(type) {
	case int:
		fmt.Println(v, "int")
	case string:
		fmt.Println(v, "string")
	default:
		fmt.Println(v, "i dont know")
	}

}

func anything(a interface{}) {
	switch v := a.(type) {
	case string:
		fmt.Println(v + "!")
	case int:
		fmt.Println(v + 1000)
	default:
		fmt.Println(v)
	}

	// switch a.(type) {
	// 	case string:
	// 		fmt.Println(a + "!")
	// 	case int:
	// 		fmt.Println(a + 1000)
	// 	default:
	// 		fmt.Println(a)
	// }
}

func f_label_for() {
Loop:
	for {
		for {
			for {
				fmt.Println("loop")
				break Loop
			}
		}
	}
	fmt.Println("end")

Loop2:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if j > 2 {
				continue Loop2
			}
			fmt.Print(j)
		}
		fmt.Println("aaa")
	}
	fmt.Println("end")

}

func f_defar() {
	defer fmt.Println("Start")
	defer fmt.Println("Middle")
	defer fmt.Println("End")
	fmt.Println("test")

	defer func() {
		fmt.Println(1)
		fmt.Println(2)
	}()

	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	file.Write([]byte("test"))
}

// panic recover
func f_panic() {
	defer func() {
		if x := recover(); x != nil {
			fmt.Println(x)
		}
	}()
	panic("runtime error")
	fmt.Println("end") // non disp
}

func f_goroutine() {
	go sub()
	go sub()

	i := 0
	for {
		i++
		fmt.Println("main loop", i)
		time.Sleep(1 * time.Second)
	}
}

func sub() {
	i := 0

	for {
		i++
		fmt.Println("sub loop", i)
		time.Sleep(1 * time.Second)
	}
}

// func init() {
// 	fmt.Println("init")
// }
