package main

import "fmt"

// slice 切片追加 和 截取
func main() {
	// （1）声明slice 是一个切片，并且初始化。长度是3 ，容量是 5
	var wanSlice = make([]int, 3, 5)

	// %v 表示打印任何的详细信息
	fmt.Printf("len = %d, cap = %d slice = %v\n", len(wanSlice), cap(wanSlice), wanSlice) // 输出： len = 3, cap = 5 slice = [0 0 0]

	// (1) 向 wanSlice 切片追加一个元素 1, wanSlice len = 4, [0,0,0,1], cap = 5
	wanSlice = append(wanSlice, 1)
	fmt.Printf("len = %d, cap = %d slice = %v\n", len(wanSlice), cap(wanSlice), wanSlice) // 输出： len = 4, cap = 5 slice = [0 0 0 1]

	// (2) 向 wanSlice 切片追加一个元素 2, wanSlice len = 5, [0,0,0,1,2], cap = 5
	wanSlice = append(wanSlice, 2)
	fmt.Printf("len = %d, cap = %d slice = %v\n", len(wanSlice), cap(wanSlice), wanSlice) // 输出： len = 5, cap = 5 slice = [0 0 0 1 2]

	// (3) 向一个容量 cap 已经满的 wanSlice 切片元素
	// 重要：len 长度增加1，cap 翻倍增加
	wanSlice = append(wanSlice, 3)
	fmt.Printf("len = %d, cap = %d slice = %v\n", len(wanSlice), cap(wanSlice), wanSlice) // 输出： len = 6, cap = 10 slice = [0 0 0 1 2 3]

	// (4) 向一个容量 cap 已经满的 wanSlice 切片元素
}
