package main

import "fmt"

func main() {

	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1: ", v1)

	fmt.Println("len:", len(m))

	delete(m, "k2")
	fmt.Println("map:", m)

	k2v, prs := m["k2"]
	_, prs2 := m["k3"]
	_, prs3 := m["k1"]
	fmt.Println("prs:", prs)
	fmt.Println("prs2:", prs2)
	fmt.Println("prs3:", prs3)
	fmt.Println("k2v:", k2v)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}
