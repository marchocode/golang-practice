package main

import (
	"errors"
	"fmt"
)

type SimpeError struct {
	message string
}

var usernameNull SimpeError = SimpeError{message: "用户名为空"}

func (s SimpeError) Error() string {
	return s.message
}

func login(u string, p string) (bool, error) {

	if u == "" {
		return false, SimpeError{message: "用户名为空"}
	}

	if p == "" {
		return false, SimpeError{message: "密码为空"}
	}

	return true, nil

}

func register(username string, password string) (int, error) {

	if "" == username {
		// 每次都产生一个新的错误，无法使用 errors.Is() 进行判断是否是同一个错误
		return 0, errors.New("用户名不能为空")
	}
	if "" == password {
		return 0, errors.New("密码不能为空")
	}

	return 1, nil
}

func main() {

	u1, err := register("", "pass1")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("u1 id=%d\n", u1)

	_, err1 := login("", "pass1")
	_, err2 := login("u2", "")

	// message信息相同，可用于判断类型
	fmt.Printf("err1 equals = %v\n", errors.Is(err1, usernameNull))
	fmt.Printf("err2 equals = %v\n", errors.Is(err2, usernameNull))

	// 转换异常，前提是必须是同一个异常类型
	var coverError SimpeError

	if errors.As(err2, &coverError) {
		fmt.Println("cover success")
		fmt.Println(coverError)
	}

}
