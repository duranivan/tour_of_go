package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["x"]++
	fmt.Println(m)

	n := []byte{126, 1, 0, 0}
	nn := string(n[:])
	fmt.Println(string(nn))

	// s := 'A'
	// fmt.Printf("%b\n", s)
	// fmt.Println(fmt.Sprintf("%v", s))
	// b := []byte("A Here is a string....")
	// fmt.Println(b)
	// fmt.Println(len(b))

	st := []byte("A")
	bytes := []byte{}
	fmt.Println(len(bytes))
	fmt.Println(bytes)
	bytes = append(bytes, st...)
	fmt.Println(bytes)
	fmt.Println(len(bytes))
	fmt.Println(st)
	fmt.Println(bytes)
}
