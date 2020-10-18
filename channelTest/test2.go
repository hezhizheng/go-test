package main

import "log"

func main() {
    ch := make(chan int, 5)

    // 写入channel 不能超过预定的容量 否则 all goroutines are asleep - deadlock!
    for i := 0; i <= 4; i++ {
        ch <- i
    }

    // 遍历读取前先close掉 ，否则 all goroutines are asleep - deadlock!
    close(ch)

    for v := range ch {
        log.Println(v, len(ch), ch)
    }

}
