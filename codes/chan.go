package main

import (
    "log"
    "sync"
    "time"
)

var wg sync.WaitGroup

/*
定义 sync
*/
func main2() {
    userCount := 10

    ch := make(chan int, 2)

    for i := 0; i < userCount; i++ {
        wg.Add(1)

        go func(i int) {
            ch <- i

            c := i + 1
            time.Sleep(time.Second)
            // 执行完之后
            wg.Done()
            d := <-ch

            log.Println("ccccccccc", c, d)

        }(i)
    }

    wg.Wait()
}

func main() {
    ch := make(chan int, 2)

    count := 10

    for i := 1; i <= count; i++ {
        wg.Add(1)

        go func(i int) {
            ch <- i

            c := i + 1
            // 模拟主程序执行处理时间
            time.Sleep(time.Second)

            d := <-ch
            wg.Done()

            log.Printf("%d + 1  = %d ，chan= %d ", i, c, d)

        }(i)
    }

    wg.Wait()
}
