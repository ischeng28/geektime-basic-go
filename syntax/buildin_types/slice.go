package main

func Slice() {
	s1 := []int{1, 2, 3, 4}
	println(s1)
	s2 := make([]int, 3, 4)
	s2 = append(s2, 7)
	s2 = append(s2, 8)
	s3 := make([]int, 4)
	println(s3)
}
