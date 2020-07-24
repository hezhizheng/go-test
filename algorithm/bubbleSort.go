package main

import "log"

func bubbleSort(ary []int) []int {
	aryLen := len(ary)

	for j := 1; j < aryLen; j++ {
		for i := 0; i < aryLen-1; i++ {
			//log.Println(i)
			if ary[i] > ary[i+1] {
				temp := ary[i+1]
				ary[i+1] = ary[i]
				ary[i] = temp
			}
		}
	}

	return ary
}

func main() {
	ary := []int{1, 2, 4, 9, 6, 5, 3}
	sort := bubbleSort(ary)

	log.Println(sort)
}
