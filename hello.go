package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}

func init() {
	bs := [3]int{1, 3, 4}
	var bd = []int{2, 3, 4}
	bd = append(bd, 3)
	fmt.Println(cap(bs))
	fmt.Println(cap(bd))

}
