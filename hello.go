package main

import "fmt"

func main2() {
	fmt.Println("Hello, World!")
}

func init2() {
	bs := [3]int{1, 3, 4}
	var bd = []int{2, 3, 4}
	bd = append(bd, 3)
	fmt.Println(cap(bs))
	fmt.Println(cap(bd))

}
