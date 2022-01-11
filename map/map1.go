package main

import "fmt"

func main() {
	// （1）声明wanMap 是一个map类型，key是string value是string
	var wanMap map[string]string

	if wanMap == nil {
		fmt.Println("wanMap is nil")
	}

	// 在使用map之前，需要先用make给map分配数据空间
	wanMap = make(map[string]string)
	wanMap["one"] = "PHP"
	wanMap["two"] = "Golang"
	wanMap["three"] = "C++"

	// 没有顺序，hash
	fmt.Println(wanMap) // 输出： map[one:PHP three:C++ two:Golang]

	// （2）第二种声明方式, 通过make开辟空间
	wanMap2 := make(map[string]string)
	fmt.Println(wanMap2) // 输出： map[]

	wanMap2["one"] = "PHP"
	wanMap2["two"] = "Golang"
	wanMap2["three"] = "C++"
	fmt.Println(wanMap2) // 输出： map[one:PHP three:C++ two:Golang]

	// （3）第三种声明方式
	// 切片slice wanSlice := []int{1,2,3,4}
	wanMap3 := map[string]string{
		"one":   "php",
		"two":   "golang",
		"three": "c++",
	}
	fmt.Println(wanMap3) // 输出： map[one:php three:c++ two:golang]

}
