package main

import (
	"fmt"
	"github.com/jianfengye/collection"
	"sync"
)

var safe threadSafety2

func main() {

	fmt.Println("123123")

	var wg sync.WaitGroup
	//var wg2 sync.WaitGroup

	//var mu sync.Mutex
	s := collection.NewIntCollection([]int{1, 2})
	s2 := collection.NewIntCollection([]int{1})

	//t := threadSafety2{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			safe.do(func() interface{} {
				return s.Append(1)
			})

			//threadSafety(func() {
			//	s.Append(1)
			//})
			//mu.Lock()
			//s.Append(1)
			//mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(s.Count())
	//fmt.Println(s)

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			safe.do(func() interface{} {
				return s2.Append(1)
			})

			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(s2.Count())

}

type threadSafety2 struct {
	mu  sync.Mutex
	res interface{}
}

var mu sync.Mutex

func threadSafety(x func()) {
	mu.Lock()
	defer mu.Unlock()
	x()
}

func (receiver *threadSafety2) do(x func() interface{}) {
	receiver.mu.Lock()
	defer receiver.mu.Unlock()
	receiver.res = x()
}
