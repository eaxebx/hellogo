package test

/*
	sync.WaitGroup
	Add(n) 计数器加n
	Done() 计数器减1
	Wait() 等待,直到为0
*/

import (
	"fmt"
	"sync"
)

func test() {
	var wg sync.WaitGroup
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func(i int) {
			fmt.Println(i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
