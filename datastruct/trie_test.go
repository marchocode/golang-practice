package datastruct

import (
	"bufio"
	"log"
	"os"
	"testing"
)

func TestNew(t *testing.T) {

	f, err := os.Open("5000-words.txt")

	if err != nil {
		log.Fatalln(err)
	}

	defer f.Close()

	scan := bufio.NewScanner(f)
	node := NewNode()

	for scan.Scan() {
		node.Insert(scan.Text())
	}

	re := node.SearchPrefix("desc", 10)
	t.Log(re)
}
