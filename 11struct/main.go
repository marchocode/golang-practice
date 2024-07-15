package main

import (
	"fmt"
)

type User struct {
	id   int
	name string
}

func (user *User) getId() {
	fmt.Println(user.id)
}

func (user *User) setId(id int) {
	user.id = id
}

// 函数的参数传递是值传递，直接传递一个对象，则会复制一份
func changeId(user User) {
	user.id = 2
}

func changeIdPoint(user *User) {
	user.id = 3
}



func main() {

	user := User{id: 1, name: "mrc"}
	fmt.Println(user)

	// 直接传递会复制
	changeId(user)
	fmt.Println(user)

	// 一般通过指针进行传递
	changeIdPoint(&user)
	fmt.Println(user)

	u2 := &User{id: 2, name: "demo"}
	// 调用结构体的方法
	u2.getId()
	u2.setId(3)

	fmt.Println(*u2)
}
