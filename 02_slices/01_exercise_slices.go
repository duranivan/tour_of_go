package main

import (
	"golang.org/x/tour/pic"
)

// Pic generates a slice of slices of uint8
func Pic(dx, dy int) [][]uint8 {
	myImage := make([][]uint8, dy)
	for i := range myImage {
		myImage[i] = make([]uint8, dx)
		for j := range myImage[i] {
			// myImage[i][j] = uint8(i + j/2)
			// myImage[i][j] = uint8(i ^ j)
			myImage[i][j] = uint8(i * j)
		}
	}

	return myImage
}

func main() {
	pic.Show(Pic)
	// myPic := Pic(256, 256)
	// fmt.Println(myPic)
	// fmt.Printf("%T\n", myPic)
}
