// Implement a rot13Reader that implements io.Reader and reads from an
// io.Reader, modifying the stream by applying the rot13 substitution cipher to
// all alphabetical characters.

// The rot13Reader type is provided for you. Make it an io.Reader by
// implementing its Read method.

package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

// TODO: use p.Read(b) inside this Read function in order to obtain b from p and transform it using the rot13 cipher.
func (p rot13Reader) Read(b []byte) (int, error) {
	n, err := p.r.Read(b)
	for i, elem := range b {
		b[i] = rot13(string(elem))
	}
	return n, err
}

func rot13(s string) byte {
	input := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	output := "NOPQRSTUVWXYZABCDEFGHIJKLMnopqrstuvwxyzabcdefghijklm"
	match := strings.Index(input, s)
	if match == -1 {
		return []byte(s)[0]
	}
	return output[match]
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	fmt.Println(s)
	r := rot13Reader{s}
	fmt.Println(r)
	io.Copy(os.Stdout, &r)
}
