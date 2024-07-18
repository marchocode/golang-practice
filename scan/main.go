package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

func connect(t string, ip string, port int, w *sync.WaitGroup) {

	defer w.Done()

	address := fmt.Sprintf("%s:%d", ip, port)
	con, err := net.DialTimeout(t, address, 2*time.Second)

	if err != nil {
		return
	}

	defer con.Close()

	fmt.Printf("Connect success %s:%d\n", ip, port)

}

func generagePort(r string) ([]int, error) {

	// 1-1024
	match, err := regexp.Match("\\d-\\d", []byte(r))

	if err != nil {
		return nil, err
	}

	if !match {
		return nil, errors.New("port range error")
	}

	strs := strings.Split(r, "-")

	start, _ := strconv.Atoi(strs[0])
	end, _ := strconv.Atoi(strs[1])

	len := (end - start) + 1

	if len < 0 {
		return nil, errors.New("port range error")
	}

	port := make([]int, len)

	for i := 0; i < (end-start)+1; i++ {
		port[i] = start + i
	}

	return port, nil
}

func main() {

	s := flag.String("s", "", "scan host for example: 1.1.1.1")
	r := flag.String("r", "1-1024", "the range of the port.")
	t := flag.String("t", "tcp", "network type: tcp or udp")

	flag.Parse()

	if flag.NFlag() == 0 {
		fmt.Println("usage scan -s [IP] -r [1-1024] -t [tcp|udp]")
		flag.PrintDefaults()
		os.Exit(1)
	}

	if *s == "" {
		fmt.Println("scan host is empty.")
		flag.PrintDefaults()
		os.Exit(1)
	}

	ports, err := generagePort(*r)

	if err != nil {
		log.Fatalln(err)
	}

	var group sync.WaitGroup
	group.Add(len(ports))

	for _, p := range ports {
		go connect(*t, *s, p, &group)
	}

	group.Wait()

}
