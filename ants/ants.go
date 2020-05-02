package main

import (
	"fmt"
	"github.com/panjf2000/ants/v2"
	"log"
	"sync"
	"time"
)

func println()  {
	log.Println("this is test ants task 1")
	//time.Sleep(time.Second)
}

func task(pool *ants.Pool) {

    for i := 1; i <= 3000; i++ {
		wg.Add(1)
		_ = pool.Submit(func() {
			println()
			wg.Done()
		})
    }
	//wg.Wait()
}

func task33() {
	for i := 1; i <= 3000; i++ {
		println()
	}
}

var (
    wg sync.WaitGroup
)

func main() {
	star := time.Now()
    // go go go
    defer ants.Release()

    pool, _ := ants.NewPool(50)
    pool2, _ := ants.NewPool(3000)

    action := func() {
        task(pool2)
		//task33()
        wg.Done()
    }

    for i := 1; i <= 50; i++ {
        wg.Add(1)
        _ = pool.Submit(action)
    }

    wg.Wait()

	endTime := time.Since(star)
    // 这里是携程执行完之后的操作

    fmt.Println("执行完毕",endTime.Seconds())

}
