package sections

import "fmt"

type User struct {
	Name string
	Age  int
}

func (u *User) SetName(name string) {
	u.Name = name
}

func Sec11() {
	fmt.Println("sec11")
	fmt.Println("--struct--")
	f_struct()
	fmt.Println("--slice-struct--")
	f_slice_struct()
}

func f_struct() {
	fmt.Println("struct")
	var user1 User
	fmt.Println(user1)
	user1.Name = "user1"
	user1.Age = 10
	fmt.Println(user1)

	//error
	user2 := User{"a", 1}
	fmt.Println(user2)

	user3 := User{Name: "user3"}
	fmt.Println(user3)

	user4 := user3
	fmt.Println(user4)
	user4.Name = "user4"
	fmt.Println(user4)
	fmt.Println(user3)
	user5 := &user3
	fmt.Println(user5)
	user5.Name = "user5"
	fmt.Println(user5)
	fmt.Println(user3)

	user6 := User{}
	UpdateUser(&user6, "user6", 20)
	fmt.Println(user6)
	user6.SetName("user6-2")
	fmt.Println(user6)

	t := T{User: User{Name: "user7", Age: 30}}
	fmt.Println(t)
	t.User.SetName("user7-2")
	fmt.Println(t)
}

func UpdateUser(user *User, name string, age int) {
	user.Name = name
	user.Age = age
}

type T struct {
	User User
}

func NewUser(name string, age int) *User {
	return &User{Name: name, Age: age}
}

type Users []*User

func f_slice_struct() {
	users := Users{}
	user1 := User{Name: "user1", Age: 10}
	user2 := User{Name: "user2", Age: 20}
	users = append(users, &user1, &user2)
	for _, user := range users {
		fmt.Println(*user)
	}
}

type Myint int

func (mi Myint) Print() {
	fmt.Println(mi)
}

func f_unique_type() {
	var mi Myint
	fmt.Println(mi)
	fmt.Printf("%T/n", mi)
	mi.Print()
}
