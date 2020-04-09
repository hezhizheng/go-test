package main

import (
    "encoding/json"
    "fmt"
    "github.com/panjf2000/ants/v2"
    "github.com/valyala/fasthttp"
    "log"
    "net/http"
    "strconv"
    "strings"
    "sync"
    "time"
)

var WaitGroup sync.WaitGroup

func main6() {
    url := `http://ltd.test/time`
    for i := 0; i < 10; i++ {
        go getTest(url, i)
    }
    select {}
}

func main61() {

    url := `http://ltd.test/time`

    ch := make(chan int, 3) //一个缓冲区大小为5的channel

    counter := 0
    for i := 0; i < 10; i++ {
        WaitGroup.Add(1) //每创建一个goroutine，就把任务队列中任务的数量+1
        go func() {
            defer WaitGroup.Done() //任务完成，将任务队列中的任务数量-1，其实.Done就是.Add(-1)
            ch <- 1

            getTest(url, i)
            counter++
            <-ch

        }()
    }

    //close(ch)
    WaitGroup.Wait()
    fmt.Println("counter:", counter)
}

func main() {

    t := time.Now()
    fmt.Println("开始:", t)
    defer ants.Release()
    antPool, _ := ants.NewPool(8)

    url := `http://ltd.test/time?id=88888888`
    //url := `https://baidu.com`
    //url := `https://hzz.cool`

    counter := 0
    for i := 0; i < 10; i++ {
        WaitGroup.Add(1) //每创建一个goroutine，就把任务队列中任务的数量+1

        antPool.Submit(func() {
            getTest(url, i)
            counter++
            WaitGroup.Done() //任务完成，将任务队列中的任务数量-1，其实.Done就是.Add(-1)
        })
    }

    go getTestBaiDu("http://ltd.test/time?id=7777777777")

    WaitGroup.Wait()

    fmt.Println("counter222:", counter)

    end := time.Since(t) // => 秒

    string2 := strconv.FormatFloat(end.Seconds(), 'f', -1, 64)

    name1 := make(map[string]interface{})
    //var name1 map[string]interface{}

    name1["rrrrrrr"] = time.Now()
    dingTalkRobot(name1)
    //SendDingMsg("qweqweqweqwe")
    fmt.Println("结束:", time.Now(), string2+"秒", end.Minutes(), end.Hours())
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
    time.Sleep(time.Second * 2)
    return string(resp)
}

func getTestBaiDu(url string) bool {

    antPool, _ := ants.NewPool(8)

    counter := 0
    for i := 0; i < 10; i++ {
        WaitGroup.Add(1) //每创建一个goroutine，就把任务队列中任务的数量+1

        antPool.Submit(func() {
            status, resp, err := fasthttp.Get(nil, url)
            if err != nil {
                fmt.Println("请求失败:", err.Error(), string(resp))
            }

            if status != fasthttp.StatusOK {
                fmt.Println("请求没有成功:", status, string(resp))
            }

            fmt.Println(string(resp), i)
            time.Sleep(time.Second * 3)
            counter++
            WaitGroup.Done() //任务完成，将任务队列中的任务数量-1，其实.Done就是.Add(-1)
        })
    }

    WaitGroup.Wait()
    fmt.Println("counter333:", counter)

    return true
}

func dingTalkRobot(msg map[string]interface{}) {
    //请求地址模板
    // "https://oapi.dingtalk.com/robot/send?access_token=" . $access_token;
    webHook := `https://oapi.dingtalk.com/robot/send?access_token=78face7560afa1524da82f63ca3fc647f5e16755c94f4e6b42f9d143081b8893`

    // $data = [
    //            'msgtype' => 'markdown',
    //            'markdown' => [
    //                'title' => $title,
    //                'text' => date('Y-m-d H:i:s') . "：\n" . "### " . $title . "\n" . $message,
    //            ]
    //        ];

    msgStr, _ := json.MarshalIndent(msg, "", " ")

    //markdown["title"] = "test"
    //markdown["text"] = "test"

    //content["msgtype"] = "markdown"
    //content["markdown"] = string(msgStr)

   //content2 := `{"msgtype": "markdown",
   //    "markdown": { "title": "测试" , "text": "` + string(msgStr) + `" }
   //}`
    var r http.Request

    r.ParseForm()
    r.Form.Add("uuid", orderUUID)
    bodystr := strings.TrimSpace(r.Form.Encode())

    content2 := `{"msgtype": "text",
		"text": {"content": "` + string(msgStr) + `"}
	}`


    //contentS, _ := json.MarshalIndent(content, "", " ")
    log.Println(content2,string(msgStr))

    //创建一个请求
    req, err := http.NewRequest("POST", webHook, strings.NewReader(content2))
    if err != nil {
        // handle error
        log.Println(err)
        return
    }

    client := &http.Client{}
    //设置请求头
    req.Header.Set("Content-Type", "application/json; charset=utf-8")
    //发送请求
    resp, err := client.Do(req)

    log.Println(resp)
    //关闭请求
    //defer resp.Body.Close()

    if err != nil {
        // handle error
        log.Println(err, resp)
        return
    }
}

func SendDingMsg(msg string) {
    //请求地址模板
    webHook := `https://oapi.dingtalk.com/robot/send?access_token=78face7560afa1524da82f63ca3fc647f5e16755c94f4e6b42f9d143081b8893`
    content := `{"msgtype": "text",
		"text": {"content": "` + msg + `"}
	}`
    //创建一个请求
    req, err := http.NewRequest("POST", webHook, strings.NewReader(content))
    if err != nil {
        // handle error
    }

    client := &http.Client{}
    //设置请求头
    req.Header.Set("Content-Type", "application/json; charset=utf-8")
    //发送请求
    resp, err := client.Do(req)
    //关闭请求
    defer resp.Body.Close()

    if err != nil {
        // handle error
    }
}
