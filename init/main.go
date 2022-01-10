package main

import (
	"golang-study/init/lib1"
	// 匿名函数导入，只会执行内部的 init() 方法
	_ "golang-study/init/lib1"

	// 别名导入
	lib11 "golang-study/init/lib1"
	"golang-study/init/lib2"
)

func main() {
	lib1.Test()
	lib2.Test()
	lib11.Test()
}
