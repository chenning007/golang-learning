package mytypes

import "fmt"

func init() {

}

func Mytype() {

	// 自定义类型
	type user struct {
		name       string
		email      string
		ext        int
		privileged bool
	}

	var lisa user
	fmt.Println(lisa)

	// 使用结构字面量创建结构类型值
	bill := user{
		name:       "bill",
		email:      "bill@email.com",
		ext:        0,
		privileged: false,
	}
	fmt.Println(bill)

	// 不使用字段名，创建结构类型的值
	bill1 := user{"bill1", "bill1@email.com", 0, true}
	fmt.Println(bill1)

	// 使用其它结构
	type admin struct {
		person user
		level  int
	}

	fred := admin{
		person: user{
			name:       "fred",
			email:      "fred@email.com",
			ext:        0,
			privileged: false,
		},
		level: 0,
	}
	fmt.Println(fred)

	type Duration int64
	var dur Duration
	// dur = int64(20) // Error in compilation
	dur = Duration(20)
	fmt.Println(dur)

}
