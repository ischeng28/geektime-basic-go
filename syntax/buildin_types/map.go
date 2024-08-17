package main

func Map() {
	m1 := map[string]int{
		"key1": 213,
	}
	m1["hello"] = 345

	m2 := make(map[string]int, 12)
	m2["key"] = 33

	val, ok := m1["hello"]
	if ok {
		println(val)
	}
}
