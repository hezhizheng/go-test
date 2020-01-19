package main

import (
	"fmt"
	"strconv"
)

func main() {
	var1 := 666
	var2 := "[]"
	ary1 := [...]int{3, 5, 6,} // 可以省略长度而采用 `...` 的方式，Go 会自动根据元素个数来计算长度
	//ary2 := [1]int{4, 5, 6,6} // 可以省略长度而采用 `...` 的方式，Go 会自动根据元素个数来计算长度

	b := []map[string]interface{}{
		{"name": "Jack", "sex": 2},
		{"name": "Mary", "sex": 1},
		{"name": "Jane", "sex": 1},
	}

	categoryId, err := strconv.Atoi("")

	if err == nil{
	}

	fmt.Println(categoryId)

	m2 := map[string]string{"a":"1a", "b":"2bb", "c":"3ccc"}

	type Html []interface{}

	html := make(Html, 5)

	html[0] = "div"

	html[1] = "span"

	html[2] = []byte("script")

	html[3] = "style"

	html[4] = "head"

	// c := []map[string]interface{}{"name": "Jack", "sex": 2}

	m3 := make(map[string]string)

	m3["sds"] = "asdasda"

	if m3["xxx"] == "" {
		fmt.Println(789)
	}

	fmt.Println(m3)
	fmt.Println(m3["xxx"])
	fmt.Printf("go funck!!!! %v asdasd %v \n",var1,var2)
	fmt.Printf("数组 %v asdasd \n",ary1)
	fmt.Printf("多维数组 %v asdasd \n",b)


	for key,value := range m2 {
		fmt.Println(key,":",value)
	}


}
