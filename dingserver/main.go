package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

var ding bool = true

func main() {

	logger := log.Default()

	http.HandleFunc("/change", func(w http.ResponseWriter, r *http.Request) {

		logger.Println("/change")
		ding = !ding
		io.WriteString(w, fmt.Sprintf("Current Status = %v\n", ding))
	})

	http.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		logger.Println("/get")
		io.WriteString(w, fmt.Sprintf("%v\n", ding))
	})

	logger.Println("Run Server On 8080")
	http.ListenAndServe(":8080", nil)
}
