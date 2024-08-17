package main

import "github.com/ischeng28/basic-go/syntax/variable/demo"

func main() {
	var a int = 123
	println(a)

	a = 234

	var b = 234
	println(b)

	var a1 = "123"
	println(a1)

	println(a1 + "hello")

	var c = 12.4
	println(c)

	var str = "hello"
	println(str)

	var d uint = 123
	println(d)

	var e int
	println(e)

	f := 123
	println(f)

	println(demo.External)
	println(demo.Global)
}

const (
	Status0 = iota
	Status1
	Status2
	Abc

	Status5 = iota + 1
	Status6 = 6
	Status7 = iota
)

const (
	DayA = iota*12 + 13
	DayB
)

const (
	MyStatus = iota<<10 + 1
	MyStatus1
)
