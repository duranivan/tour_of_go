package main

import (
	"fmt"

	"golang.org/x/tour/reader"
)

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.

func (r MyReader) Read(b []byte) (n int, err error) {
	bLen := len(b)
	n = bLen
	for i := 0; i < bLen; i++ {
		b[i] = []byte("A")[0]
	}

	return n, err
}

func main() {
	r := MyReader{}
	b := make([]byte, 8)
	fmt.Println(b)
	fmt.Println(r.Read(b))
	fmt.Println(b)
	reader.Validate(MyReader{})
}
