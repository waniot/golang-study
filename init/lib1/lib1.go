package lib1

import "fmt"

func Test() {
	fmt.Println("lib2 Test ...")
}

func init() {
	fmt.Println("lib1 init() ...")
}
