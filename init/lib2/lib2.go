package lib2

import "fmt"

// 大写，对外开放
func Test() {
	fmt.Println("lib2 Test ...")
}

func init() {
	fmt.Println("lib1 init() ...")
}
