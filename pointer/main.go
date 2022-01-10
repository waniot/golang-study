package main

import "fmt"

func swap(pa *int, pb *int) {
	var tmp int
	tmp = *pa
	*pa = *pb
	*pb = tmp
}
func main() {
	a := 100
	b := 200

	// a b 值交换
	swap(&a, &b)

	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	fmt.Println("a &", &a)
	fmt.Println("b &", &b)

	// 定义一个p指针
	var p *int

	// 把a的地址赋值给 p的内存地址
	p = &a

	// p的地址则为 a 的内存地址
	fmt.Println("p &", p)
}
