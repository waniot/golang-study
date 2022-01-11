package main

// defer 执行顺序，就是一个堆栈。即入栈和出站操作
import "fmt"

func fun1() {
	fmt.Println("fun1 A")
}

func fun2() {
	fmt.Println("fun2 B")
}

func fun3() {
	fmt.Println("fun3 C")
}

func main() {
	defer fun1() // 第一个入栈 第三个出栈
	defer fun2() // 第二个入栈 第二个出栈
	defer fun3() // 第三个入栈 第一个出栈

	// 打印结果
	// fun3 C
	// fun2 B
	// fun1 A
}
