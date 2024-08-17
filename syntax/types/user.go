package main

import "fmt"

type User struct {
	Name     string
	Age      int
	NickName string
}

func (u User) ChangeName(name string) {
	u.Name = name
}

func (u *User) ChangeAge(age int) {
	u.Age = age
}

func ChangeUser() {
	u1 := User{Name: "Tom", Age: 18}
	fmt.Println(u1)
	fmt.Println(&u1)

	u1.ChangeName("wangc")
	u1.ChangeAge(213)
	fmt.Println(u1)

	println("---------------")

	u2 := &User{Name: "Jerry", Age: 82}
	fmt.Println(u2)
	fmt.Println(&u2)
	u2.ChangeName("wangc")
	u2.ChangeAge(23)
	fmt.Println(u2)
}

func NewUser() {
	//	初始化结构体
	u := User{}
	u.Name = "Jerry"
	fmt.Println(u.Name)

	u1 := &User{}
	fmt.Println(u1)

	u1 = new(User)
	fmt.Println(u1)

	var u3 User
	fmt.Println(u3)

	var u4 *User
	fmt.Println(u4)

	u5 := User{Name: "cheng"}
	fmt.Println(u5)

	u5 = User{Name: "ww", Age: 13, NickName: "chengc"}
	fmt.Println(u5)

}

func UseList() {
	l1 := LinkedList{}
	l1ptr := &l1
	var l2 LinkedList = *l1ptr
	fmt.Println(l2)

	var l3 *LinkedList
	fmt.Println(l3)

}
