package mycollections

import "fmt"

func MappingExamples() {

	// 创建映射
	dict1 := make(map[string]string)
	fmt.Println(dict1)
	dict2 := map[string][]string{}
	fmt.Println(dict2)
	dict3 := map[string]string{}
	fmt.Println(dict3)

	// 使用映射
	//
	dict3["a"] = "a-value"
	fmt.Println(dict3) //map[a:a-value]

	// nil dict declaration
	//var dict4 map[string]string
	//dict4["b"] = "b-value"
	// runtime error

	value, exists := dict3["a"]
	if exists {
		fmt.Println(value)
	} // a-value

	fmt.Println(dict3["a"]) // a-value

	value2 := dict3["a"]
	fmt.Println(value2) // a-value

	// 映射中删除key
	delete(dict3, "a")
	fmt.Println(dict3) //map[]

}
