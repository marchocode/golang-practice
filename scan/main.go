package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"sync"
	"time"
)

func connect(ip string, port int, w *sync.WaitGroup) {

	defer w.Done()

	address := fmt.Sprintf("%s:%d", ip, port)
	con, err := net.DialTimeout("tcp", address, 1*time.Second)

	if err != nil {
		return
	}

	defer con.Close()

	fmt.Printf("Connect success %s:%d\n", ip, port)

}

func generagePort() []int {

	port := make([]int, 1000)

	for i := 0; i < 1000; i++ {
		port[i] = i + 1
	}

	return port
}

func main() {

	host := flag.String("host", "", "your host name")

	flag.Parse()

	if flag.NArg() == 0 {
		flag.PrintDefaults()
		os.Exit(0)
	}

	fmt.Println(*host)

	/*
			 *Port numbers are assigned in various ways, based on three ranges: System
		Ports (0-1023), User Ports (1024-49151), and the Dynamic and/or Private
		Ports (49152-65535)
	*/

}
