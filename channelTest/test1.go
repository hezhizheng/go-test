package main

import (
    "fmt"
    "log"
    "time"
)

const limit  = 5000

func setNumToChan(numChan chan int) {
    for i := 1; i <= limit; i++ {
        numChan <- i
        //fmt.Println("写入numChan", i)
    }
    // 遍历前 需要先close掉
    close(numChan)
}

func calculationChan(numChan chan int, resChan chan int) {
    for v := range numChan {
        resNum := 0
        for i := 1; i <= v; i++ {
            resNum += i
        }
        resChan <- resNum
        //log.Println("当前取出numChan", v, "resChan总和", resNum)
    }
    //close(resChan)
}

func rangResChan(resChan chan int) {
    // 遍历取出channel的时候需要先close掉
    close(resChan)
    for v := range resChan {
        log.Println("rangResChan", v)
    }
}

func main() {

    //resNum := 0
    //for i := 1; i <= 3; i++ {
    //   resNum += i
    //}
    //fmt.Println(resNum)
    //time.Sleep(time.Second)
    //return


    t := time.Now()
    numChan := make(chan int, limit)
    resChan := make(chan int, limit)
    //
    go setNumToChan(numChan)

    go calculationChan(numChan, resChan)
    go calculationChan(numChan, resChan)
    go calculationChan(numChan, resChan)
    go calculationChan(numChan, resChan)
    go calculationChan(numChan, resChan)
    go calculationChan(numChan, resChan)
    go calculationChan(numChan, resChan)
    go calculationChan(numChan, resChan)

    for {
        if len(numChan) == 0 && len(resChan) == limit {

            rangResChan(resChan)

            if len(resChan) == 0 {
                end := time.Since(t)
                fmt.Println("耗时",end)
                break
            }
        }
    }

}
