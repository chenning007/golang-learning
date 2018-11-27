package mycollections

import "fmt"

func init() {

}

func SlicesExamples() {

	// 切片的内部构成（数组指针，容量，长度）

	// make与切片字面量
	slice1 := make([]string, 5) //长度和容量都是5
	fmt.Println(slice1)

	slice2 := make([]string, 3, 5) // 长度是3容量是5
	fmt.Println(slice2)

	// []为空的字变量是切片
	slice3 := []string{"str1", "str2", "str3"}
	fmt.Println(slice3)

	// 指定元素值
	slice4 := []int{3: 2, 5: 1}
	fmt.Println(slice4) //[0 0 0 2 0 1]

	// nil空切片
	var slice5 []int
	fmt.Println(slice5) //[]
	slice6 := make([]int, 0)
	fmt.Println(slice6) //[]
	slice7 := []int{}
	fmt.Println(slice7) //[]

	// 切片赋值
	slice8 := []int{1, 2, 3, 4}
	fmt.Println(slice8[0]) // 1
	slice8[0] = 5
	fmt.Println(slice8[0]) // 5
	slice9 := slice8[1:3]  // slice[i:j] 长度：j-i 容量：k-i (k是底层数组容量）
	fmt.Println(slice9)
	fmt.Println(len(slice9))
	fmt.Println(cap(slice9))
	slice8[1] = 10
	fmt.Println(slice9) // 共享数组，一个改变是影响另一个的值

	// 增长切片
	slice10 := append(slice9, 60)
	slice11 := append(slice10, 60)
	slice12 := append(slice11, 60)
	fmt.Println(slice10) // [10 3 60]
	fmt.Println(slice11) // [10 3 60 60]
	fmt.Println(slice12) // [10 3 60 60 60]

	// 超长增长切片
	slice13 := []int{1, 2, 3}
	fmt.Println("length is ", len(slice13), " and capability is ", cap(slice13))
	slice14 := append(slice13, 4)
	fmt.Println("length is ", len(slice14), " and capability is ", cap(slice14)) // double the cap

	// 3索引创建切片
	slice15 := []string{"a", "b", "c", "d", "d"}
	slice16 := slice15[2:3:4] // len = 3-2; cap = 4-2
	fmt.Println(slice16)
	fmt.Println("length is ", len(slice16), " and cap is ", cap(slice16))

	// 切片联接
	slice17 := []int{1, 2}
	slice18 := []int{3, 4}
	fmt.Println("cat result is ", append(slice17, slice18...))

	// 迭代切片
	slice19 := []int{1, 2, 3, 4}
	for index, value := range slice19 {
		fmt.Println("index %d value %d", index, value)
	}

	slice20 := []int{1, 2, 3, 4}
	for _, value := range slice20 {
		fmt.Println("value : %d, address %X", value, &value)
	}
	//value : %d, address %X 1 0xc0000181c8
	//value : %d, address %X 2 0xc0000181c8
	//value : %d, address %X 3 0xc0000181c8
	//value : %d, address %X 4 0xc0000181c8

	slice21 := []int{1, 2, 3, 4}
	for index := 2; index < len(slice21); index++ {
		fmt.Println(slice21[index])
	}

	// 多维切片
	slice22 := [][]int{{1}, {1, 2}}
	fmt.Println(slice22)

}
