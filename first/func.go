package main

import "fmt"

// (1)返回值和类型
func fun1(a string, b int) int {
	fmt.Println("a ", a)
	fmt.Println("b ", b)

	c := 100
	return c
}

// (2) 多个返回值和类型
func fun2(a string, b int) (int, string) {
	fmt.Println("a ", a)
	fmt.Println("b ", b)

	c := 100
	d := "Hi Tinywan"
	return c, d
}

// (3) 多个返回值和类型，指定返回值名称
func fun3(a string, b int) (r1 int, r2 string) {
	fmt.Println("a ", a)
	fmt.Println("b ", b)

	// 给有名称的返回值变量赋值
	r1 = 2022
	r2 = "Tinywan"
	return
}

// (4) 多个返回值和类型，指定返回值名称
func fun4(a string, b int) (r1, r2 int) {
	fmt.Println("a ", a)
	fmt.Println("b ", b)

	// 给有名称的返回值变量赋值
	r1 = 2022
	r2 = 10086
	return
}

func main() {
	fun1("Tinywan", 2022)

	res1, res2 := fun2("PHP", 20222)
	fmt.Println("res1 = ", res1, "res2 = ", res2)

	r1, r2 := fun3("Golang", 1990)
	fmt.Println("r1 = ", r1, "r2 = ", r2)
}
