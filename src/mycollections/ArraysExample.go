package mycollections

import "fmt"

func init() {

}

func ArraysExamples() {

	// 1. 数组声明和初始化
	var array [5]int
	var array0 = [5]int{}
	var array1 = [5]int{0, 1, 2, 3, 4}
	var array2 = [...]int{0, 1, 2, 3, 4}
	var array3 = [5]int{1: 1, 2: 2}
	fmt.Println(array)               //[0 0 0 0 0]
	fmt.Println(array0, len(array0)) // [0 0 0 0 0] 5
	fmt.Println(array1, len(array1)) // [0 1 2 3 4] 5
	fmt.Println(array2, len(array2)) // [0 1 2 3 4] 5
	fmt.Println(array3, len(array3)) // [0 10 20 0 0] 5

	// 2. 使用数组
	// 2.1 访问数组元素
	fmt.Println(array2[2]) // 2
	array2[2] = -2
	fmt.Println(array2[2]) // -2

	var array4 = [5]*int{0: new(int), 1: new(int)}
	fmt.Println(array4) // [0xc00006e018 0xc00006e020 <nil> <nil> <nil>]
	*array4[0] = 10
	// *array4[2] = 20 // nil exception
	*array4[0] = 20
	fmt.Println(*array4[0], array4) // 20 [0xc000018070 0xc000018078 <nil> <nil> <nil>]

	// 3. 数组复制
	var array5 [2]string
	var array6 = [2]string{"red", "blue"}
	fmt.Println(array6) // [red blue]
	fmt.Println(array5) // [ ]
	array5 = array6
	fmt.Println(array5) // [red blue]
	array5[1] = "yellow"
	fmt.Println(array5) // [red yellow]
	fmt.Println(array6) // [red blue]

	// TODO := vs = 赋值的差异
	// 指针数组
	var array7 [3]*string
	array8 := [3]*string{new(string), new(string), new(string)}
	*array8[0] = "str1"
	*array8[1] = "str2"
	*array8[2] = "str3"
	array7 = array8
	fmt.Println(array8)     //[0xc0000621d0 0xc0000621e0 0xc0000621f0]
	fmt.Println(*array7[0]) // str1
	fmt.Println(*array8[0]) // str1
	*array7[0] = "str-updated"
	fmt.Println(*array7[0]) // str-updated
	fmt.Println(*array8[0]) // str-updated

	// 多维数组
	var array9 = [2][3]int{{1, 1}, {2, 2}}
	fmt.Println(array9) // [[1 1 0] [2 2 0]]
	// TODO
	// var array10 = [2][3]int{1:{1,1}, 2:{2,2}}

	// 大数组的数据复制在函数调用时有性能问题
	// 更好的方式是通过地址传递给函数

}
