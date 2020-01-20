package main

import (
    "fmt"
    "github.com/jasonlvhit/gocron"
    "github.com/robfig/cron/v3"
    "log"
)

func main() {
    i := 0
    c := cron.New(cron.WithSeconds()) // v3 用法 干！
    // 1 3 5 * * ?  每天 05:03:01 执行
    //spec := "*/2 * * * * ?"
    spec := "0 18 16 * * ?"
    c.AddFunc(spec, func() {
        i++
        log.Println("cron running:", i)
    })
    c.Start()

    defer c.Stop()

    select{}

}

func task() {
    fmt.Println("I am running task.")
}

func main11()  {

    gocron.Every(1).Day().At("17:15").Do(task)

    <- gocron.Start()
}
