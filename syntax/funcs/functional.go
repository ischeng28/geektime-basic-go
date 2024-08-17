package main

func Fun4() {
	myFunc3 := Func3
	s, err := myFunc3(1, 2)
	println(s, err)
	_, _ = Func3(2, 3)
}

func Functional6() {
	//	新定义了一个方法，赋值给了fn
	fn := func() string {
		return "heelo"
	}
	fn()
}

// Functional8 匿名方法立即调用
func Functional8() {
	//	新定义了一个方法，赋值给了fn
	fn := func() string {
		return "heelo"
	}()
	println(fn)
}

// Functional7 返回一个 返回string的无参方法
func Functional7() func() string {
	return func() string {
		return ""
	}
}
