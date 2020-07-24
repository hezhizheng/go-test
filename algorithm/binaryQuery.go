package main

import (
    "log"
)

// 定义二分查找的接口
type BinaryQueryInterface interface {
    main() int
}

// 使用递归的方式实现接口
type Recursive struct {
    BinaryQueryInterface
}

// 普通二分查找
type General struct {
    BinaryQueryInterface
}

// 定义递归的参数
type RecursiveParam struct {
    Ary    []int
    Target int
    Low    int
    Top    interface{}
}

// 实现递归二分的方法
func (r *Recursive) main() int {
    // ary为有序数据
    ary := []int{1, 2, 4, 8, 9, 11, 13}
    // 如果需要默认值最好是通过结构体去定义
    rp := RecursiveParam{
        Ary:    ary,
        Target: 8,
        Low:    0,
        Top:    nil,
    }
    return recursive(rp)
}

// 递归二分查找算法
func recursive(param RecursiveParam) int {
    top := param.Top
    low := param.Low
    ary := param.Ary
    target := param.Target

    aryLen := len(ary)

    if param.Top == nil {
        top = aryLen
    }

    // 返回的是一个不大于mid的int值
    mid := (top.(int) + low) / 2

    isSet := inArray(aryLen, func(i int) bool {
        return target == ary[i]
    })

    if !isSet {
        return -1 // 找不到
    }

    if ary[mid] == target {
        return mid
    } else if ary[mid] > target {
        param.Top = mid - 1
        return recursive(param)
    } else {
        param.Low = mid + 1
        return recursive(param)
    }

}

// 实现普通二分查找的方法
func (g *General) main() int {
    ary := []int{1, 2, 4, 5, 6, 7, 8}
    return general(ary, 71)
}

// 普通二分查找的算法
func general(ary []int, target int) int {
    aryLen := len(ary)
    top := aryLen
    low := 0

    mid := -1
    for low <= top {
        mid = (top + low) / 2
        isSet := inArray(aryLen, func(i int) bool {
            return target == ary[i]
        })

        if !isSet {
            return -1 // 找不到
        }

        if ary[mid] == target {
            return mid
        } else if ary[mid] > target {
            top = mid - 1
        } else {
            low = mid + 1
        }
    }
    return mid
}

//
func inArray(aryLen int, f func(i int) bool) bool {
    for i := 0; i < aryLen; i++ {
        if f(i) {
            return true
        }
    }
    return false
}

func main() {
    // 实现递归接口
    c := BinaryQueryInterface.main(&Recursive{})
    log.Printf("通过递归查找的数组角标为 %d", c)

    c2 := BinaryQueryInterface.main(&General{})
    log.Printf("普通二分查找的数组角标为 %d", c2)
}
