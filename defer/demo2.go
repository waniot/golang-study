package main

// defer 和 return 谁先后：return 先执行，defer 后执行
import "fmt"

func deferFunc() int {
	fmt.Println("defer func called ...")
	return 0
}

func rerunFunc() int {
	fmt.Println("return func called ...")
	return 0
}

func deferAndReturn() int {
	defer deferFunc()
	return rerunFunc()
}

func main() {
	defer deferAndReturn()

	// 打印结果
	// return func called ...
	// defer func called ...
}
