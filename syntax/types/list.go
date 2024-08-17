package main

type List interface {
	Add(index int, val any)
	Append(val any) error
	Delete(index int) error
}

type LinkedList struct {
	head node
}

type node struct {
	next *node
}

func (l *LinkedList) Append(val any) error {
	panic("implement me")
}

func (l *LinkedList) Add(index int, val any) {

}

func (l *LinkedList) Delete(index int) error {
	panic("implement me")
}

func UserListV1() {
	l := &LinkedList{}
	l.Add(1, 213)
	l.Add(1, "123")
	l.Add(1, nil)
}
