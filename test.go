package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["x"]++
	fmt.Println(m)

	n := []byte{126, 1, 0, 0}
	nn := string(n[:])
	fmt.Println(string(nn))
}
