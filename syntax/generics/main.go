package main

// T 类型参数 名字叫做T 约束是any 等于没有约束
type List[T any] interface {
	Add(idx int, t int)
	Append(t int)
}

func UseList() {
	var l List[int]
	l.Append(213)

}
