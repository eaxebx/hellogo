package main

import (
	"fmt"
	_ "hellogo/testdatabase"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(50)
	wg.Add(50)
	wg.Add(1)
	for i := 0; i < 100; i++ {
		go func(i int) {
			fmt.Println(i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
