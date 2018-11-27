package mytypes

import (
	"fmt"
)

func init() {

}

type notifier interface {
	notify1()
	notify2()
}

func (u *user) notify1() {
	fmt.Println("send email to ", u.name, " , his email is ", u.email)
}

func (u user) notify2() {
	fmt.Println("send email to ", u.name, " , his email is ", u.email)
}

func sendNotification(n notifier) {
	n.notify1()
	n.notify2()
}

func MyInterfaces() {

	bill := user{"bill", "bill@email.com"}
	sendNotification(&bill)

}
