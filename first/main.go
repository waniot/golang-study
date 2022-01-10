package main

import (
	"fmt"
	"time"
)

var gA int = 100

// gB := 200 这样子不支持全局变量

// 多变量声明
var gC, gD = 200, 300
var (
	gE int  = 400
	gF bool = false
)

// 定义常量
const BEIJING = 10086
const SHANGHAI = 10000

func main() {
	time.Sleep(1 * time.Second)
	fmt.Println("Hello World")

	a := 1
	fmt.Printf("a type is = %T\n", a)

	fmt.Println("gA = ", gA)
	fmt.Println("gC = ", gC)
	fmt.Println("gD = ", gD)
	fmt.Println("gE = ", gE)
	fmt.Println("gF = ", gF)
}
