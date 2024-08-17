package main

func Closure(name string) func() string {
	//闭包
	//name变量
	//方法本身
	return func() string {
		return "hello," + name
	}
}
