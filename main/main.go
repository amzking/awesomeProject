package main

import "fmt"


type user struct {
	name string
	email string
}

func (u user) notify() {
	fmt.Printf("sending %s: <%s>\n", u.name, u.email)
}

func (u * user) changeEmail(email string) {
	u.email = email
}

func (u user) changeEmail1(email string) {
	u.email = email
}


func main() {

	i := 3;
	fmt.Println(&i);
	fmt.Println("test")

	bill := user {"Bill", "bill@email.com"}
	bill.notify();

	lisa := &user{"Lisa", "lisa@email.com"}
	lisa.notify();

	bill.changeEmail("bill@email.cn")
	bill.notify();

	lisa.changeEmail("lisa@email.cn")
	lisa.notify()
}
