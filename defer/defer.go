package main

import "fmt"

func main() {
	// 函数结束之前执行的函数defer， PHP析构函数

	defer fmt.Println("main 1 end")
	fmt.Println("main go 1")

	defer fmt.Println("main 2 end")
	fmt.Println("main go 2")
}
