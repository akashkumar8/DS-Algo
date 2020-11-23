package main

import "fmt"

func main() {

	m := make(map[string]int)

	m["a1"] = 3
	m["a2"] = 11

	fmt.Println("map:", m)

	v1 := m["a1"]
	fmt.Println("v1: ", v1)

	fmt.Println("len:", len(m))

	delete(m, "a2")
	fmt.Println("map:", m)

	_, prs := m["a2"]
	fmt.Println("prs:", prs)

	n := map[string]int{"cafe": 1, "bar": 2}
	fmt.Println("map:", n)
}
