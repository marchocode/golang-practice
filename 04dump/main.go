package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	scan := bufio.NewScanner(os.Stdin)
	m := make(map[string]int)

	for scan.Scan() {

		text := scan.Text()

		if text == "/" {
			break
		}
		m[text]++
	}

	for k, v := range m {
		fmt.Println("k=", k, "value=", v)
	}

}
