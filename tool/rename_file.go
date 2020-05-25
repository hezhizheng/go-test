package main

import (
    "io/ioutil"
    "log"
    "os"
    "strings"
    "sync"
    "time"
)

var wg = sync.WaitGroup{}

func main() {

    star := time.Now()

    pathname := "C:\\E\\music"

    rd, err := ioutil.ReadDir(pathname)

    if err != nil {
        panic(err)
    }

    channel := make(chan int, 10)

    for _, f := range rd {
        wg.Add(1)
        fileName := f.Name()
        //log.Println(fileName, pathname+"\\"+fileName)

        go rename(pathname, fileName, channel)

    }

    wg.Wait()

    end := time.Since(star)

    log.Printf("任务完成，耗时 %f 秒", end.Seconds())

}

// 陈百强 - 等.mp3 重命名为 陈百强-等.mp3
func rename(path, f string, channel chan int) {
    channel <- 1
    // 重命名文件
    file := f

    split := strings.Split(file, " - ")

    // 大于1 才认为是需要重命名的
    if len(split) > 1 {
        //log.Println(split)

        oldPath := path + "\\" + file

        newPath := path + "\\" + split[0] + "-" + split[1]

        err1 := os.Rename(oldPath, newPath)
        if err1 != nil {
            panic(err1)
        } else {
            log.Printf("文件 %s 重命名为=> %s ", oldPath, newPath)
        }
    }

    // 模拟并发
    //time.Sleep(time.Second * 1)
    <-channel
    wg.Done()
}
