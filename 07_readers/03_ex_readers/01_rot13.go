package main

import (
	"bytes"
	"fmt"
)

func rot13(x ...byte) []byte {
	input := []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
	output := []byte("NOPQRSTUVWXYZABCDEFGHIJKLMnopqrstuvwxyzabcdefghijklm")
	result := []byte{}
	for _, e := range x {
		match := bytes.Index(input, []byte{e})
		if match == -1 {
			result = append(result, e)	
		} else {
			result = append(result, output[match])
		}
	}
	return result
}

func main() {
	myS :=  []byte("Uryyb Jbeyq! ab srne")
	fmt.Println(myS)
	fmt.Println(rot13(myS...))
	fmt.Println(string(rot13(myS...)))
	fmt.Printf("%T\n", rot13(myS...))
}
