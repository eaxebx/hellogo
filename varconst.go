package main

import "fmt"

func varstest() {

	// 3种方式
	var name string
	var age = 16
	bnb := true

	// 多个变量声明
	var a, b, c int
	a, b, c = 1, 2, 3

	var e, f, g = 4, 5, 6

	x, y, z := 10, 20, 30

	var (
		m int
		n string
	)
	name = "Alice"
	fmt.Println(name, age, bnb, a, b, c, e, f, g, x, y, z, m, n)

}
