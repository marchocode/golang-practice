package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func simpleRead() {

	file, err := os.Open("go.mod")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()

	/*
		*for {

			buf := make([]byte, 32)
			n, err := file.Read(buf)

			if err != nil {
				fmt.Println(err)
				break
			}

			fmt.Printf("read num =%d \n", n)
			fmt.Println(string(buf[:n]))

		}
	*/

	// or
	bytes, err := io.ReadAll(file)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("read num = %d cap= %d\n", len(bytes), cap(bytes))
	fmt.Println(string(bytes))

}

func bufferRead() {

	file, err := os.Open("go.mod")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	buf := bufio.NewReader(file)
	bf := make([]byte, 32)

	for {

		n, err := buf.Read(bf)

		if err != nil && err == io.EOF {
			fmt.Println("over")
			break
		}

		fmt.Printf("read num =%d \n", n)
		fmt.Println(string(bf[:n]))
	}

}

func main() {

	bufferRead()
}
