package main

import (
	"fmt"
	"net"
)

func read(con net.Conn) {

	defer con.Close()

	defer func() {
		fmt.Println("client close.")
	}()

	buf := make([]byte, 1024)

	for {

		n, err := con.Read(buf)

		if err != nil {
			return
		}

		fmt.Printf("receive number %d \n", n)
		fmt.Println(string(buf[:n]))
	}

}

func main() {

	listen, err := net.Listen("tcp", "127.0.0.1:8080")

	if err != nil {
		fmt.Printf("error.\n")
	}

	fmt.Println("listen to 127.0.0.1:8080")

	for {

		con, err := listen.Accept()

		if err != nil {
			fmt.Printf("receive error.\n")
		}

		fmt.Printf("new connection %s.\n", con.RemoteAddr().String())

		go read(con)
	}

}
