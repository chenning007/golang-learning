package mytypes

import "fmt"

func init() {

}

type user struct {
	name  string
	email string
}

func (u user) notify() {
	fmt.Println("send email to user ", u.email)
}

func (u *user) changeEmail(newEmail string) {
	u.email = newEmail
}

func Myfunc() {
	bill := user{"bill", "bill@email.com"}
	bill.notify() // send email to user  bill@email.com
	(&bill).changeEmail("bill@hotmail.com")
	bill.notify() // send email to user  bill@hotmail.com
	bill.changeEmail("bill@sina.com.cn")
	bill.notify() // send email to user  bill@sina.com.cn
}
