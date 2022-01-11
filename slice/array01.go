package main

import "fmt"

func printArrayLen(wanArray [4]int) {
	for index, value := range wanArray {
		fmt.Println("printArrayLen index = ", index, "value = ", value)
	}
	wanArray[0] = 111
}
func main() {
	// (1) 固定长度的数组，默认值都是 0
	var wanArray1 [10]int

	wanArray2 := [10]int{1, 2, 3, 4, 5, 6}
	wanArray3 := [4]int{1, 2, 3, 4}

	for i := 0; i < len(wanArray1); i++ {
		fmt.Println(wanArray1[i])
	}

	// range 遍历数组和切片
	for index, value := range wanArray2 {
		fmt.Println("index = ", index, "value = ", value)
	}

	// 查看数组的数据类型
	fmt.Printf("wanArray1 types = %T\n", wanArray1) // wanArray1 types = [10]int
	fmt.Printf("wanArray2 types = %T\n", wanArray2)
	fmt.Printf("wanArray3 types = %T\n", wanArray3) // wanArray3 types = [4]int

	// 调用固定长度的数组：副本，值拷贝
	printArrayLen(wanArray3)
	// 尝试打印修改数组长度
	for index, value := range wanArray3 {
		fmt.Println("index = ", index, "value = ", value)
	}

	// （重要）调用[动态]数组
	printArrayLen(wanArray3)
}
