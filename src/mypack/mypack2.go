package mypack

import "fmt"

// 同一个包下不能定义同名的函数，即使在不同的文件中也不行
func Myprint2() {
	fmt.Println("this is print from myprint2")
}
