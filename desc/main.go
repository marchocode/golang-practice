package main

import (
	"fmt"
	"net"
)

func main() {

	con, err := net.Dial("tcp", "182.254.198.208:20014")

	if err != nil {
		fmt.Println("connect error.")
		return
	}

	defer con.Close()

	for {

		buf := make([]byte, 0, 1024)
		n, err := con.Read(buf)

		if err != nil {
			fmt.Printf("recvice num = %d\n", n)
		}
	}

}
