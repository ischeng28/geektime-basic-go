package main

import (
	"math"
	"strconv"
	"strings"
	"unicode/utf8"
)

func main() {
	var a int = 456
	var b int = 123
	println(a + b)
	println(a - b)
	println(a * b)
	println(a / b)
	println(float64(a) / float64(b))
	//	a=a+1
	a++
	println(a)
	//	b=b-1
	b--
	println(b)

	//编译不通过
	//var c float64 = 12.3
	//println(a + c)

	// 编译不通过
	//var d int32 = 12
	//println(a + d)

	println(math.Abs(-12.3))

	ExtremeNum()

	String()
	Byte()
	Bool()
}

func ExtremeNum() {
	println(math.MinInt64)
	println("float64 最小正数", math.SmallestNonzeroFloat64)
	println("float32 最小正数", math.SmallestNonzeroFloat32)
}

func String() {
	//	He said: "Hello Go!"
	println("He said: \"Hello Go!\"")
	println(`hello,go
换行了
再一行
	结束hello go`)
	println(strconv.Itoa(123))
	println(len("hello"))
	println(len("hello你好"))
	println(utf8.RuneCountInString("hello你好"))

	println(strings.CutPrefix("hello你好", "h"))
	println(strings.CutPrefix("hello你好", "E"))

}

func Byte() {
	var a byte = 'a'
	println(a)

	var str string = "hello"
	var bs []byte = []byte(str)
	var str1 string = string(bs)
	println(str1)
}

func Bool() {
	var a bool = true
	var b bool = false
	println(a && b)
	println(a || b)
	println(!a)
}
