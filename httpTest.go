package main

import (
    "fmt"
    "github.com/panjf2000/ants/v2"
    "github.com/valyala/fasthttp"
    "sync"
)

var WaitGroup sync.WaitGroup

func main1() {

    url := `http://ltd.test/time`

    ch := make(chan int, 1) //一个缓冲区大小为5的channel

    for i := 0; i < 100000; i++ {
        WaitGroup.Add(1) //每创建一个goroutine，就把任务队列中任务的数量+1
        go func() {
            defer WaitGroup.Done() //任务完成，将任务队列中的任务数量-1，其实.Done就是.Add(-1)
            ch <- 1

            getTest(url, i)
            <-ch
        }()
    }

    //close(ch)
    WaitGroup.Wait()
}

func main() {
    defer ants.Release()
    antPool, _ := ants.NewPool(10)

    //url := `http://ltd.test/time`
    //url := `https://baidu.com`
    url := `https://hzz.cool`

    for i := 0; i < 100000; i++ {
        WaitGroup.Add(1) //每创建一个goroutine，就把任务队列中任务的数量+1

        antPool.Submit(func() {
            getTest(url, i)
            WaitGroup.Done() //任务完成，将任务队列中的任务数量-1，其实.Done就是.Add(-1)
        })
    }

    WaitGroup.Wait()
}

func getTest(url string, i int) string {

    status, resp, err := fasthttp.Get(nil, url)
    if err != nil {
        fmt.Println("请求失败:", err.Error(), string(resp))
    }

    if status != fasthttp.StatusOK {
        fmt.Println("请求没有成功:", status, string(resp))
    }

    fmt.Println(string(resp), i)
    //time.Sleep(time.Second * 2)
    return string(resp)
}
