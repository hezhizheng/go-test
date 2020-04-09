package main

import (
    "fmt"
    "regexp"
)

func main() {
    str := "第01集$http://down.okdown7.com/20200401/18506_91a8fa3b/花甲男孩转大人EP01.mp4#第02集$http://down.okdown7.com/20200401/18507_001e5ecd/花甲男孩转大人EP02.mp4#第03集$http://down.okdown7.com/20200401/18508_206a2ef3/花甲男孩转大人EP03.mp4#第04集$http://down.okdown7.com/20200401/18509_32238672/花甲男孩转大人EP04.mp4#第05集$http://down.okdown7.com/20200401/18510_5b393ed8/花甲男孩转大人EP05.mp4#第06集$http://down.okdown7.com/20200401/18511_2250ed6f/花甲男孩转大人EP06.mp4#第07集$http://down.okdown7.com/20200401/18512_5a677bbc/花甲男孩转大人EP07.mp4"
    matched, err := regexp.MatchString("https?://([\\w-]+\\.)+[\\w-]+(/[\\w-./?%&=]*)?", str)
    fmt.Println(matched, err)

    r, _ := regexp.Compile("https?://([\\w-]+\\.)+[\\w-]+(/[\\w-./?%&=]*)?")

    mp4 := r.FindAllString(str, -1)

    fmt.Println(mp4, err)
}
