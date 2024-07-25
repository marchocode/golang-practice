package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id       int    `json:id`
	Name     string `json:name`
	Password string `json:password`
}

func main() {

	t := reflect.TypeOf(User{})

	for i := 0; i < t.NumField(); i++ {

		f := t.Field(i)
		fmt.Printf("name =%s tag=%v\n", f.Name, f.Tag)

	}

	sli1 := []string{"a", "b", "c"}
	sli2 := []string{"a", "b", "c"}
	sli3 := []string{"d", "e", "f"}

	// 判断两个string切片是否相等
	fmt.Println("sli1 == sli2 ? ", reflect.DeepEqual(sli1, sli2))
	fmt.Println("sli1 == sli3 ? ", reflect.DeepEqual(sli1, sli3))

	u := User{Id: 1, Name: "test", Password: "123"}
	fmt.Printf("%%v = %v\n", u)
	fmt.Printf("%%+v = %+v\n", u)
	fmt.Printf("%%#v = %#v\n", u)

}
