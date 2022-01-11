package main

import "fmt"

func main() {
	// （1）声明slice 是一个切片，并且初始化，默认值是 1,2,3,4 长度len是4
	wanSlice := []int{1, 2, 3, 4}

	// %v 表示打印任何的详细信息
	fmt.Printf("len = %d, slice = %v\n", len(wanSlice), wanSlice) // 输入为： len = 4, slice = [1 2 3 4]

	// （2）声明slice 是一个切片，但是没有给slice分配空间
	var wanSlice2 []int
	fmt.Printf("wanSlice2 len = %d, slice = %v\n", len(wanSlice2), wanSlice2) // 输出： wanSlice2 len = 3, slice = []
	// 通过make 来给slice 开辟空间
	wanSlice2 = make([]int, 3)
	fmt.Printf("wanSlice2 len = %d, slice = %v\n", len(wanSlice2), wanSlice2) // 输出： wanSlice2 len = 3, slice = [0 0 0]

	// 给开辟空间赋值
	wanSlice2[0] = 10086
	wanSlice2[1] = 10000
	fmt.Printf("wanSlice2 len = %d, slice = %v\n", len(wanSlice2), wanSlice2) // 输出： wanSlice2 len = 3, slice = [10086 10000 0]

	// （3）声明slice 是一个切片，同时给slice分配空间,初始化值是 0
	var wanSlice3 []int = make([]int, 3)
	fmt.Printf("wanSlice3 len = %d, slice = %v\n", len(wanSlice3), wanSlice3) // 输出： wanSlice3 len = 3, slice = [0 0 0]

	wanSlice31 := make([]int, 3)
	fmt.Printf("wanSlice31 len = %d, slice = %v\n", len(wanSlice31), wanSlice31) // 输出： wanSlice31 len = 3, slice = [0 0 0]

	// （4）判断slice是否已经开辟了空间
	var wanSlice4 []int
	if wanSlice4 == nil {
		fmt.Println("wanSlice4 是一个空切片") // 输出： wanSlice4 是一个空切片
	} else {
		fmt.Println("wanSlice4 是有空间的")
	}
}
