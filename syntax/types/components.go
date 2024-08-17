package main

type Outer struct {
	Inner
}

type Inner struct {
}

type Outer1 struct {
	*Inner
}

func (i Inner) Hello() {
	println("hello innner")
}

func Components() {
	var o Outer
	o.Hello()

	o1 := Outer1{
		Inner: &Inner{},
	}

	o1.Hello()
}
