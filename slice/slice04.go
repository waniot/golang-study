package main

import "fmt"

// slice 截取
func main() {
	s := []int{1, 2, 3}

	// %v 表示打印任何的详细信息
	fmt.Printf("len = %d, cap = %d slice = %v\n", len(s), cap(s), s) // 输出： len = 3, cap = 3 slice = [1 2 3]

	// （1）截取
	s1 := s[0:2]
	fmt.Println(s1) // 输出： [1 2]

}
