package main

import (
	"fmt"
	"sync"
)

type SafeSlice struct {
	sync.Mutex
	s []int
}

func (ss *SafeSlice) Append(val int) {
	ss.Lock()
	ss.s = append(ss.s, val)
	ss.Unlock()
}

func (ss *SafeSlice) Len() int {
	ss.Lock()
	defer ss.Unlock()
	return len(ss.s)
}

func main() {
	var wg sync.WaitGroup
	ss := &SafeSlice{s: make([]int, 0)}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			ss.Append(4)
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(ss.Len())
}
