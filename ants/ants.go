package main

import (
	"fmt"
	"github.com/panjf2000/ants/v2"
	"sync"
)

func task() {
	fmt.Println("this is test ants task 1")
	//time.Sleep(time.Second)
}

var (
	wg sync.WaitGroup
)

func main() {
	// go go go
	defer ants.Release()

	pool, _ := ants.NewPool(5)

	action := func() {
		task()
		wg.Done()
	}

	for i := 1; i <= 100; i++ {
		wg.Add(1)
		_ = pool.Submit(action)
	}

	wg.Wait()

	// 这里是携程执行完之后的操作

	fmt.Println("执行完毕")

}
