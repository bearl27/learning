package main

import "fmt"

// 型が違うが同様な扱いをしたい
type Stringfy interface {
	ToString() string
}

type Person struct {
	Name string
	Age  int
}

func (p *Person) ToString() string {
	return fmt.Sprintf("Name=%v, Age=%v", p.Name, p.Age)
}

type Car struct {
	Number string
	Model  string
}

func (c *Car) ToString() string {
	return fmt.Sprintf("Name=%v, Age=%v", c.Number, c.Model)
}

func sec12() {
	vs := []Stringfy{
		&Person{Name: "Bob", Age: 2},
		&Car{Number: "123-456", Model: "Toyota"},
	}
	for _, v := range vs {
		fmt.Println(v.ToString())
	}

	err := RaiseError()
	fmt.Println(err.Error())
	e, ok := err.(*MyError)
	if ok {
		fmt.Println(e.ErrCode)
		fmt.Println(e.Aho)
	}
	f_stringer()
}

type MyError struct {
	Message string
	ErrCode int
	Aho     string
}

func (e *MyError) Error() string {
	return e.Message
}
func RaiseError() error {
	return &MyError{Message: "エラーが発生しました", ErrCode: 500, Aho: "あほ"}
}

// type Stringer interface {
// 	String() string
// }

type Point struct {
	X int
	Y string
}

func (p *Point) String() string {
	return fmt.Sprintf("X=%v, Y=%v", p.X, p.Y)
}

func f_stringer() {
	p := &Point{100, "hoge"}
	fmt.Println(p)
}
