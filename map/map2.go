package main

import "fmt"

func printMap(wanMap map[string]string) {
	// wanMap 是一个引用传递
	for key, value := range wanMap {
		fmt.Println("printMap , key =", key, " value =", value)
	}
}

// 尝试修改map值
func updateMap(wanMap map[string]string) {
	// wanMap 是一个引用传递
	wanMap["Nickname"] = "ShaoBoWan"
}

func main() {
	wanMap := make(map[string]string)

	fmt.Println(wanMap) // 输出 map[]

	// （1）添加
	wanMap["China"] = "Beijing"
	wanMap["USA"] = "NewWork"
	wanMap["Japan"] = "Tokyo"
	fmt.Println(wanMap) // 输出 map[China:Beijing Japan:Tokyo USA:NewWork]

	// （2）遍历
	for key, value := range wanMap {
		fmt.Println("key =", key, " value =", value)
	}

	// （3）删除
	delete(wanMap, "China")
	fmt.Println(wanMap) // 输出  map[Japan:Tokyo USA:NewWork]

	// （4）修改
	wanMap["USA"] = "Tinywan"
	fmt.Println(wanMap) // 输出  map[Japan:Tokyo USA:Tinywan]

	printMap(wanMap)

	updateMap(wanMap)
	fmt.Println(wanMap) // 输出  map[Japan:Tokyo Nickname:ShaoBoWan USA:Tinywan]
}
