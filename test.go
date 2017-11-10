package main

import (
	"bytes"
	"fmt"
)

func rot13(x byte) byte {
	input := []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
	output := []byte("NOPQRSTUVWXYZABCDEFGHIJKLMnopqrstuvwxyzabcdefghijklm")
	match := bytes.Index(input, []byte{x})
	if match == -1 {
		return x
	}
	return output[match]
}

func main() {
	var myS byte = 65
	fmt.Println(rot13(myS))
	fmt.Println(string(rot13(myS)))
	fmt.Printf("%T\n", rot13(myS))

	myString := []byte("Lbh penpxrq gur pbqr!")
	fmt.Println(myString)
	for i, elem := range myString {
		myString[i] = rot13(elem)
	}
	fmt.Printf("%s\n", myString)
}
