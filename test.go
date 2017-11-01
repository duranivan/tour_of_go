package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["x"]++
	fmt.Println(m)
}
