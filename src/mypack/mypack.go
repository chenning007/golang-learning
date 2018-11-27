// 包名与所在目录名一到
// 将项目路径设置到GOPATH里
package mypack

import "fmt"

// init方法在main调用之前被调用
func init() {
	fmt.Println("this is mypack init")
}

// 方法名以大写字母开头，包外才可以被访问
func Myprint(param string) {
	fmt.Println("this is myprint output: ", param)
}
