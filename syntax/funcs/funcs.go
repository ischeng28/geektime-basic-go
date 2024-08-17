package main

// Func0 单一返回值
func Func0(name string) string {
	return "hello world"
}

// Func1 多个返回值
func Func1(a, b, c int, str string) (string, error) {
	return "", nil
}

// Func2 带名字的返回值
func Func2(a int, b int) (str string, err error) {
	str = "hello"
	//带名字的返回值 可以直接return
	return
}

// Func3 带名字的返回值
func Func3(a int, b int) (str string, err error) {
	res := "hello"
	//虽然带名字 但是没有用
	return res, nil
}

// Func4 多个参数，一个类型
func Func4(a, b string) {

}
func Func5(a string, b string) {

}

func Func6(a string, b string) string {
	//	有返回值就一定要返回
	return ""
}

func Func7(a string, b string) (string, string) {
	return "hello", "world"
}

func Func8() (name string, age int) {
	return "cheng", 18
}

func Func9() (name string, age int) {
	name = "cheng"
	age = 18

	return
}

func Func10() (name string, age int) {
	//等价于 "",0
	return
}
