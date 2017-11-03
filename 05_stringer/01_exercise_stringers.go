package main

import (
	"fmt"
	"strings"
)

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.

// String returns a formmated IPAddr as "1.2.3.4"
func (i IPAddr) String() string {
	myIPAddr := make([]string, 4)
	for j, elem := range i {
		myIPAddr[j] = fmt.Sprintf("%v", elem)
	}

	myString := strings.Join(myIPAddr, ".")

	return myString
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		// fmt.Printf("%v: %v\n", name, ip.String())
		// fmt.Printf("%v: %v\n", name, ip)
		fmt.Printf("%v: %q\n", name, ip)
	}
}
