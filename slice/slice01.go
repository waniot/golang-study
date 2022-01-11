package main

import "fmt"

// []int 切片类型，这里的传值方式为 地址应用传递动态数组
func printArray(wanArray []int) {
	for _, value := range wanArray {
		fmt.Println("value = ", value)
	}

	// 应用传递：即指针
	wanArray[0] = 100
	wanArray[1] = 200
}

func main() {
	// 数组长度为空 [] 表示是一个动态数组
	wanArray := []int{1, 2, 3, 4}

	fmt.Printf("wanArray type is %T\n", wanArray) // 输入为： []int 表示为 切片类型

	printArray(wanArray)

	for _, value := range wanArray {
		fmt.Println("value = ", value)
	}
	// slice 是一个动态数组
}
