package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("123123")

	//var mu sync.Mutex
	var wg sync.WaitGroup
	s := []int{1, 2, 3}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			//mu.Lock()
			s = append(s, 1)
			//mu.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(len(s))
	fmt.Println(s)
}
