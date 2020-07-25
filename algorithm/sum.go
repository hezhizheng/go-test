package main

import "log"

// 两数之和 todo 待调整
func sum(ary []int, target int) []int {

    aryLen := len(ary)
    indexs := []int{}

    for key,value := range ary{
        index := target - value
        log.Println(index,key)
        isSet := inArray2(aryLen, func(i int) bool {
            return ary[i] == index
        })

        if isSet && index!= ary[key] {
            indexs = append(indexs,key)
        }
    }

    return indexs
}

func inArray2(aryLen int, f func(i int) bool) bool {
    for i := 0; i < aryLen; i++ {
        if f(i) {
            return true
        }
    }
    return false
}

func main() {
    sum := sum([]int{1, 2, 4, 5, 6, 7}, 3)
    log.Println(2, sum)
}
